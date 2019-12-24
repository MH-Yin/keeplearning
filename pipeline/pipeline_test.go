package pipeline

import (
	"fmt"
	"testing"
)

func Test_pipline(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	o1 := sq(done, sq(done, gen(2, 3)))
	o2 := sq(done, sq(done, gen(5, 6)))

	for re := range merge(done, o1, o2) {
		fmt.Println(re)
	}
}
