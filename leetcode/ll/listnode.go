package ll

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	if ln == nil {
		return "nil"
	}
	var s = fmt.Sprint(ln.Val)
	var node = ln.Next
	for node != nil {
		s += "->" + fmt.Sprint(node.Val)
		node = node.Next
	}
	return s
}

func initList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	var node = &ListNode{
		Val: vals[0],
	}
	var it = node
	for i := 1; i < len(vals); i++ {
		it.Next = &ListNode{
			Val: vals[i],
		}
		it = it.Next
	}
	return node
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var res *ListNode
	var next = head
	var pre *ListNode

	for next != nil {
		pre = res
		res = next
		next = next.Next
		res.Next = pre
	}
	return res
}
