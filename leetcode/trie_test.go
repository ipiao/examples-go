package leetcode

import (
	"testing"

	_261 "github.com/ipiao/examples-go/leetcode/1261"
)

func TestTrie(t *testing.T) {
	book := []string{"WordsFrequency", "get", "geta", "gete", "get", "get", "i", "have", "a", "apple"}
	wf := _261.Constructor(book)
	t.Log(wf)

	t.Log("geta==", wf.Get("geta"))
	t.Log("gt==", wf.Get("gt"))
}
