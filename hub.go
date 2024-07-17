package broadcast

import "sync"

type Hub[T any] struct {
	mu          sync.RWMutex
	subscribers map[chan T]struct{}
}

func (h *Hub[T]) Broadcast(msg T) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for ch := range h.subscribers {
		select {
		case ch <- msg:
		default: // Drop msg if listener is not listening
		}
	}
}

func (h *Hub[T]) Subscribe() Subscriber[T] {
	ch := make(chan T)

	h.mu.Lock()
	defer h.mu.Unlock()
	if h.subscribers == nil {
		h.subscribers = make(map[chan T]struct{})
	}
	h.subscribers[ch] = struct{}{}
	return Subscriber[T]{
		C:   ch,
		hub: h,
	}
}

func (h *Hub[T]) unsubscribe(c chan T) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.subscribers, c)
}
