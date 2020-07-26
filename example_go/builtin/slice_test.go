package builtin

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {

	ss := make([]int, 0, 8)
	ss = append(ss, 1, 2, 3, 4, 5, 6)
	cs := ss
	nss := append(ss, 7)
	ncs := append(cs, 8)
	t.Log(ss, cs)
	t.Log(nss, ncs)
}

func TestSliceSubSets(t *testing.T) {
	subsets([]int{9, 0, 3, 5, 7})
}

func subsets(nums []int) [][]int {
	var ret [][]int
	if len(nums) == 0 {
		return ret
	}
	ret = append(ret, []int{})
	for _, num := range nums {
		fmt.Println("new append:", num)
		l := len(ret)
		for i := 0; i < l; i++ {
			// nis:=make([]int,len(is))
			// copy(nis,is)
			is := ret[i]
			fmt.Println("is before append:", is)
			nis := append(is, num)
			fmt.Println("append nis:", nis)
			ret = append(ret, nis)
			fmt.Println("is after append:", is)
			fmt.Println()
		}
		fmt.Println("===================")
	}
	return ret
}

func TestSliceRangeAppend(t *testing.T) {
	a := []int{1, 2, 3, 4}
	for i := range a {
		a = append(a, i)
	}
	t.Log(a)
}

func l(a []int) int {
	fmt.Println("innnn")
	return len(a)
}
