package main

// 常用递归，上层开始往下层调用
func Fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return Fib(n-1) + (n - 2)
}

// n, Fib(n+1), Fib(n+2)
func TailFib(n int, ret1, ret2 int) int {
	if n == 1 {
		return ret1
	}
	return TailFib(n-1, ret2, ret1+ret2)
}

//
func main() {
	//f:== Fib(10)
	//tf:=Fib()
}
