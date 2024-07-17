package broadcast

type Subscriber[T any] struct {
	C   chan T
	hub *Hub[T]
}

func (s Subscriber[T]) Close() {
	s.hub.unsubscribe(s.C)
}
