package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var slow = head
	var fast = head
	var pre *ListNode
	for fast.Next != nil && fast.Next.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	node := &TreeNode{
		Val: slow.Val,
	}
	if pre != nil {
		pre.Next = nil
		node.Left = sortedListToBST(head)
	}

	node.Right = sortedListToBST(slow.Next)
	return node
}
