package builtin

import (
	"encoding/binary"
	"testing"
)

func TestBin(t *testing.T) {
	var i uint = 1
	t.Logf("%0b", i)
	var b = make([]byte, 64)
	binary.BigEndian.PutUint64(b, 1)
	t.Log(b)
}
