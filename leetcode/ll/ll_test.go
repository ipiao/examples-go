package ll

import "testing"

func TestReverse(t *testing.T) {
	ln := initList([]int{1, 2, 3, 4, 5})
	t.Log(ln)
	res := reverseList(ln)
	t.Log(res)
}
