package example_alg

import "testing"

// 常用递归，上层开始往下层调用
func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return fib(n-1) + (n - 2)
}

// go test -v -run 'TestRecursive'
func TestRecursive(t *testing.T) {
	_ = fib(10)
}
