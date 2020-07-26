package main

import (
	"fmt"
)

func main() {
	fmt.Println(smallestK([]int{1, 3, 5, 7, 2, 4, 6, 8}, 4))
}

func smallestK(arr []int, k int) []int {
	var h = make([]int, 0, k+1)
	// init heap
	for i := 0; i < len(arr); i++ {
		h = append(h, arr[i])
		upHeap(h, len(h)-1)
		if i >= k {
			h[0], h[k] = h[k], h[0]
			downHeap(h, 0, k)
			h = h[:k]
		}
	}
	return h
}

// 最小K个数
// 最大顶堆
func upHeap(heap []int, j int) {
	for {
		i := (j - 1) / 2
		if i == j || heap[i] > heap[j] {
			break
		}
		heap[i], heap[j] = heap[j], heap[i]
		j = i
	}
}

// 山区堆中最大的数
func downHeap(heap []int, i0, n int) {
	i := i0
	for {
		j := 2*i + 1
		if j >= n {
			break
		}
		if j2 := j + 1; j2 < n && heap[j2] > heap[j] {
			j = j2
		}
		if heap[i] > heap[j] {
			break
		}
		heap[i], heap[j] = heap[j], heap[i]
		i = j
	}
}
