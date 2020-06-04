package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	min := k * (k + 1) / 2
	max := 45
	if n < min || n > max {
		return nil
	}
	res := generate(k, n, 1)
	return res
}

// k个数，和为n，最小值为min
func generate(k int, n int, min int) [][]int {
	if k == 1 {
		if n >= min && n <= 9 {
			return [][]int{{n}}
		}
		return nil
	}
	if (min+min+k-1)*k/2 > n {
		return nil
	}
	g1 := generate(k-1, n-min, min+1)
	g2 := generate(k, n, min+1)
	for _, g := range g1 {
		g = append(g, min)
		g2 = append(g2, g)
	}
	return g2
}

func main() {
	res := combinationSum3(3, 15)
	fmt.Println(res)
}
