package container

import (
	"container/heap"
	"container/ring"
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	r := ring.New(5)
	i := 1
	r.Value = 1

	next := r.Next()
	for next != r {
		i++
		next.Value = i
		next = next.Next()
	}

	r.Do(func(i interface{}) {
		t.Log(i)
	})
}

func TestHeap(t *testing.T) {
	h := &MonthHeap{"Jan", "Feb", "Mar"}
	heap.Init(h)
	heap.Push(h, "May")
	heap.Push(h, "Apr")
	// first month: Jan
	fmt.Println("first month:", (*h)[0])

	// 输出: Jan	Feb	Mar	Apr	May
	for h.Len() > 0 {
		fmt.Printf("%s\t", heap.Pop(h)) // 注意不是h.Pop()
	}
	fmt.Println()
}
