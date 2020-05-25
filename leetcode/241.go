package leetcode

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(input string) []int {
	nums, ops := parse(input)
	// 迭代
	var l = len(nums)
	var ress = make([][][]int, l) // i,j,res,表示从i到j数字能够组合出的所有情况
	for i := range nums {
		ress[i] = make([][]int, l)
		ress[i][i] = []int{nums[i]}
	}

	// 表达式包含的数量多少
	for n := 2; n <= l; n++ {
		for i := 0; i < l-n+1; i++ {
			j := i + n - 1
			var res []int
			// 分成2部分运算
			for k := i; k < j; k++ {
				r1 := ress[i][k]
				r2 := ress[k+1][j]

				for _, n1 := range r1 {
					for _, n2 := range r2 {
						r := cal(n1, n2, ops[k])
						res = append(res, r)
					}
				}
			}
			ress[i][j] = res
		}
	}

	fmt.Println(ress)
	return ress[0][l-1]
}

// 解析字符串
func parse(s string) ([]int, []rune) {
	var numstack []int
	var operstack []rune

	start := 0
	for i, r := range s {
		switch r {
		case '+', '-', '*':
			numstack = append(numstack, getNum(s[start:i]))
			operstack = append(operstack, r)
			start = i + 1
		}
	}
	numstack = append(numstack, getNum(s[start:]))
	return numstack, operstack
}

func getNum(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func cal(num1, num2 int, op rune) int {
	var r int
	switch op {
	case '+':
		r = num1 + num2
	case '-':
		r = num1 - num2
	case '*':
		r = num1 * num2
	}
	return r
}
