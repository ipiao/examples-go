package builtin

import "testing"

func TestForLen(t *testing.T) {
	ss := []int{1, 2, 3}
	for len(ss) > 0 {
		t.Log(ss[0])
		ss = ss[1:]
	}
}
