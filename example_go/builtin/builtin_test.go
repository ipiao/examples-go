package builtin

import (
	"fmt"
	"testing"
)

func TestUint0(t *testing.T) {
	fmt.Printf("%x", ^uint(0)>>63)
}
