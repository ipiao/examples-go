package ll

func ringLength(node *ListNode) int {
	// 首先确定有环
	fast := node
	slow := node
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			couner := 1
			for slow.Next != fast {
				couner++
				slow = slow.Next
			}
			return couner
		}
	}
	return 0
}
