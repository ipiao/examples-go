package builtin

import (
	"fmt"
	"testing"
)

func TestUint0(t *testing.T) {
	fmt.Printf("%x", ^uint(0)>>63)
}

func TestCopy(t *testing.T) {
	b := make([]byte, 0, 10)
	copy(b, "hello")
	t.Log(string(b))
}
