package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type bitmap []byte

func (bm *bitmap) set(target int) {
	b := *bm
	g := target / 8
	i := target % 8
	if g+1 > len(*bm) {
		for i := 0; i < g+1-len(*bm); i++ {
			b = append(b, 0)
		}
	}
	b[g] |= 1 << i
	*bm = b
}

func (bm *bitmap) has(target int) bool {
	g := target / 8
	i := target % 8
	if g+1 > len(*bm) {
		return false
	}
	return (*bm)[g]&(1<<i) != 0
}

type FindElements struct {
	node *TreeNode
	bm   *bitmap
}

func Constructor(root *TreeNode) FindElements {
	bm := make(bitmap, 0)
	ele := FindElements{
		node: root,
		bm:   &bm,
	}
	generate(0, root, ele.bm)
	return ele
}

func generate(val int, node *TreeNode, bm *bitmap) {
	node.Val = val
	bm.set(val)
	if node.Left != nil {
		generate(val*2+1, node.Left, bm)
	}
	if node.Right != nil {
		generate(val*2+2, node.Right, bm)
	}
}

func (this *FindElements) Find(target int) bool {
	return this.bm.has(target)
}

func main() {
	node := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: -1,
						Left: &TreeNode{
							Val:   -1,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:   6,
				Right: nil,
			},
		},
	}

	ele := Constructor(node)
	log.Println(ele.Find(6))
	log.Println(ele.Find(9))
	log.Println(ele.Find(4))
	log.Println(ele.Find(10))
	log.Println(ele.Find(9))
	for _, b := range *ele.bm {
		log.Printf("%b", b)
	}

}
