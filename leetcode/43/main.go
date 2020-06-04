package main

import "fmt"

// 0-9
func charToInt(c byte) int {
	return int(c - '0')
}

func intToChar(i int) byte {
	return byte('0' + i)
}

func multiply(num1 string, num2 string) string {
	//for i:=0; i<len(num1)
	//
	//
	//return ""
}

func main() {
	fmt.Println(charToInt('9') == 9)
	fmt.Println(intToChar(9) == '9')
}
