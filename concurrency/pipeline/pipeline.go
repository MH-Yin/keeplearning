package pipeline

import (
	"sync"
)

func gen(num ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range num {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for b := range in {
			select {
			case out <- b * b:
			case <-done:
				return
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(ch <-chan int) {
		defer wg.Done()
		for n := range ch {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
