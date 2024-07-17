package broadcast_test

import (
	"fmt"
	"sync"
	"time"

	"github.com/cyrusaf/broadcast"
)

func ExampleHub() {
	hub := broadcast.Hub[string]{}

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Create new subscriber to hub
			s := hub.Subscribe()
			defer s.Close() // Close subscriber when its no longer needed
			msg := <-s.C    // Listen for messages broadcast over the hub
			fmt.Println(msg)
		}()
	}

	time.Sleep(100 * time.Millisecond)
	hub.Broadcast("Hello World!")
	wg.Wait()

	//Output:
	//Hello World!
	//Hello World!
	//Hello World!
	//Hello World!
	//Hello World!
}
