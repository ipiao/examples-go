package encoding

import (
	"encoding/binary"
	"testing"
)

func TestBinary(t *testing.T) {
	type S struct {
		A int64
		B bool
	}

	t.Log(binary.Size(int64(0)))   // -1
	t.Log(binary.Size(true))       // -1
	t.Log(binary.Size(S{A: 1024})) // -1

	b := []byte{'k', 1, 1, 1, 1, 1}
	t.Log(binary.Uvarint(b))
}
