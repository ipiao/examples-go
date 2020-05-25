package leetcode

import "fmt"

func numTilePossibilities(tiles string) int {
	var nums = make(map[byte]int)
	var sum int
	for i := range tiles {
		nums[tiles[i]]++
	}
	var ta []int
	for _, i := range nums {
		ta = append(ta, i)
	}

	var bta = make([]int, len(ta))
	copy(bta, ta)
	for {
		if sumTa(bta) > 0 {
			sum += allComp(bta)
			fmt.Println(bta, sum)
			bta = generate(ta, bta)
		} else {
			break
		}
	}

	return sum
}

func generate(a []int, b []int) []int {
	for j := len(b) - 1; j >= 0; j-- {
		if b[j] >= 1 {
			b[j]--
			break
		} else {
			b[j] = a[j]
		}
	}
	return b
}

func sumTa(a []int) int {
	var sum int
	for _, t := range a {
		sum += t
	}
	return sum
}

// 所有的组合情况
func allComp(arr []int) int {
	sum := 0
	for _, a := range arr {
		sum += a
	}
	ret := ss(sum)
	for _, a := range arr {
		ret /= ss(a)
	}

	return ret
}

func ss(n int) int {
	ret := 1
	for i := 1; i <= n; i++ {
		ret *= i
	}
	return ret
}
