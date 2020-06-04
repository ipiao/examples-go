package main

import "fmt"

func permutation(S string) []string {
	cps := allComp([]byte(S))
	var res []string
	for _, cp := range cps {
		res = append(res, string(cp))
	}
	return res
}

func allComp(b []byte) [][]byte {
	if len(b) == 1 {
		return [][]byte{b}
	}
	// 去重
	var hasmap = make(map[int]bool)
	// 因为长度在1-9
	var valmap = make(map[byte]int)
	for i, lb := range b {
		if _, ok := valmap[lb]; !ok {
			valmap[lb] = i + 1
		}
	}

	var res [][]byte
	cps := allComp(b[1:])
	for _, cp := range cps {
		fmt.Println("==================\n", string(cp))
		for i := 0; i <= len(cp); i++ {
			ncp := make([]byte, len(cp)+1)
			copy(ncp, cp[:i])
			ncp[i] = b[0]
			copy(ncp[i+1:], cp[i:])
			fmt.Println(string(ncp))
			if cv, has := hasComp(valmap, hasmap, ncp); !has {
				hasmap[cv] = true
				res = append(res, ncp)
			} else {
				fmt.Println("has ", string(ncp))
			}
		}
	}
	return res
}

func hasComp(vm map[byte]int, hasmap map[int]bool, comp []byte) (int, bool) {
	var compVal int
	for i, b := range comp {
		compVal += vm[b] * (1 << (4 * i))
	}
	return compVal, hasmap[compVal]
}

func main() {
	res := permutation("Frx")
	fmt.Println(res)
}
