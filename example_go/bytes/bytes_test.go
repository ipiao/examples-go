package bytes

import (
	"bytes"
	"testing"
)

func TestBytesCompare(t *testing.T) {
	a := make([]byte, 5, 10)
	b := make([]byte, 5, 8)
	t.Log(bytes.Compare(a, b))   // zero
	t.Log(bytes.Equal(a, b))     // true
	t.Log(bytes.EqualFold(a, b)) // true
}

func TestBytesRune(t *testing.T) {
	s := "中文"
	b := []byte(s)
	t.Log(b)
	r := bytes.Runes(b)
	t.Log(r)
	i := bytes.IndexFunc(b, func(r rune) bool {
		return r == '文'
	})
	t.Log(i)
}

func TestBytesCount(t *testing.T) {
	s := []byte("aa")
	ss := []byte("aaa")
	n := bytes.Count(ss, s)
	t.Log(n)
}

func TestBytes(t *testing.T) {
	s := []byte("abcde")
	i := bytes.IndexAny(s, "ed")
	t.Log(i)

	s1 := []byte("this is a title")
	st := bytes.Title(s1)
	t.Log(string(st))
	stt := bytes.ToTitle(s1)
	t.Log(string(stt))
}

func TestBytesToString(t *testing.T){
	s:="abc"
	b:=[]byte(s)
	b[1]='d'
	t.Log(s)
}
