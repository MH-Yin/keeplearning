package context

import (
	"fmt"
	"sync"
	"testing"
)

func Test_oldModelTask(t *testing.T) {
	var wg sync.WaitGroup
	done := make(chan interface{})
	defer close(done)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(done); err != nil {
			fmt.Println(err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(done); err != nil {
			fmt.Println(err)
			return
		}
	}()

	wg.Wait()
}
