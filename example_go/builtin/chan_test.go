package builtin

import "testing"

func TestChan(t *testing.T) {
	var ch = make(chan int, 3)
	close(ch)
}

func closeCh(ch <-chan int) {
	close(ch)
}
