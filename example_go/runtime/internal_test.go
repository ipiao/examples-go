package runtime

import (
	"fmt"
)

type Uintreg uint64

// stubs: 存根，存根是用来验证票据并留底的
func ExampleStuds() {
	fmt.Println(4 << (^uintptr(0) >> 63))
	fmt.Println(4 << (^Uintreg(0) >> 63))
	// Output:
	// 8
	// 8
}
