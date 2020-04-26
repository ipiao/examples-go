package leetcode

import "encoding/json"

type WordsFrequency struct {
	trie *WordTrie
}

type WordTrie struct {
	Word     string             // 前缀 or word
	Count    int                //
	Children map[byte]*WordTrie //
}

func (wt WordsFrequency) String() string {
	b, _ := json.Marshal(wt.trie)
	return string(b)
}

func getWordFromTrie(trie *WordTrie, word string, prefix byte) *WordTrie {
	child := trie.Children[prefix]
	if child != nil {
		if len(word) == len(child.Word) {
			if word == child.Word {
				return child
			}
		} else {
			if len(word) > len(child.Word) {
				if word[:len(child.Word)] == child.Word {
					return getWordFromTrie(child, word, word[len(child.Word)])
				}
			}
		}
	}
	return nil
}

func addWordToTrie(trie *WordTrie, word string, prefix byte) {
	if trie.Children[prefix] == nil {
		trie.Children[prefix] = &WordTrie{
			Word:     word,
			Count:    1,
			Children: make(map[byte]*WordTrie),
		}
	} else {
		child := trie.Children[prefix]

		lw := len(word)
		lcw := len(child.Word)

		minlen := lw
		if minlen > lcw {
			minlen = lcw
		}
		var sublen int // 子串长度
		for i := 0; i < minlen; i++ {
			if word[i] == child.Word[i] {
				sublen++
			} else {
				break
			}
		}

		if sublen == lcw {
			if sublen < len(word) {
				addWordToTrie(child, word, word[sublen])
			} else {
				child.Count++
			}
		} else {
			subword := word[:sublen]
			wt := &WordTrie{
				Word:     subword,
				Count:    0,
				Children: make(map[byte]*WordTrie),
			}
			trie.Children[prefix] = wt
			nprefix := child.Word[sublen]
			wt.Children[nprefix] = child

			if sublen < lw {
				wwt := &WordTrie{
					Word:     word,
					Count:    1,
					Children: make(map[byte]*WordTrie),
				}
				wnprefix := word[sublen]
				wt.Children[wnprefix] = wwt
			} else {
				wt.Count += 1
			}
		}
	}
}

func Constructor(book []string) WordsFrequency {
	wf := WordsFrequency{
		trie: &WordTrie{
			Children: make(map[byte]*WordTrie),
		},
	}
	for _, word := range book {
		addWordToTrie(wf.trie, word, word[0])
	}
	return wf
}

func (this *WordsFrequency) Get(word string) int {
	subTrie := getWordFromTrie(this.trie, word, word[0])
	if subTrie != nil {
		return subTrie.Count
	}
	return 0
}

/**
 * Your WordsFrequency object will be instantiated and called as such:
 * obj := Constructor(book);
 * param_1 := obj.Get(word);
 */
