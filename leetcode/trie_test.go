package leetcode

import "testing"

func TestTrie(t *testing.T) {
	book := []string{"WordsFrequency", "get", "geta", "gete", "get", "get", "i", "have", "a", "apple"}
	wf := Constructor(book)
	t.Log(wf)

	t.Log("geta==", wf.Get("geta"))
	t.Log("gt==", wf.Get("gt"))
}
