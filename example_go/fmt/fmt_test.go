package fmt

import (
	"fmt"
	"testing"
)

type S struct {
	A  string
	I  int
	T  bool
	PA *string
}

func TestFmt(t *testing.T) {
	s := new(S)

	fmt.Printf("%v\n", s)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%#v\n", s)
	fmt.Printf("%T\n", s)

	fmt.Printf("%b\n", 9)
	fmt.Printf("%c\n", 65)
	fmt.Printf("%o\n", 65)
	fmt.Printf("%q\n", 65) // ASCII码 , A
	fmt.Printf("%U\n", 65)

}
