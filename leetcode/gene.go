package leetcode

func generateMatrix(n int) [][]int {
	//
	num := 0
	var i, j int
	c := 0
	var a = make([][]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n, n)
	}
	for num < n*n-1 {
		for j = c; j < n-1-c; j++ {
			num++
			a[c][j] = num
		}
		for i = c; i < n-c-1; i++ {
			num++
			a[i][n-c-1] = num
		}
		for j = n - c - 1; j > c; j-- {
			num++
			a[n-c-1][j] = num
		}
		for i = n - c - 1; i > c; i-- {
			num++
			a[i][c] = num
		}
		c++
	}
	if n%2 == 1 {
		a[n/2][n/2] = n * n
	}
	return a
}
