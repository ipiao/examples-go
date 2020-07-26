package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	l := len(matrix)
	r := len(matrix[0])

	var i, j int
	for {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			return false
		}
		if i < l-1 && matrix[i+1][j] <= target && (j >= r-1 || matrix[i+1][j] <= matrix[i][j+1]) {
			i = i + 1
			continue
		} else if j < r-1 && matrix[i][j+1] <= target {
			j = j + 1
			continue
		}
		return false
	}
}

func main() {
	r := findNumberIn2DArray([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30}}, 5)
	fmt.Println(r)
}
