package builtin

import (
	"fmt"
	"testing"
)

func TestGofn(t *testing.T) {
	//runtime.GOMAXPROCS(1)
	//
	//for i := 1; i <= 10; i++ {
	//	go func(n int) {
	//		log.Println(n)
	//	}(i)
	//}
	//time.Sleep(time.Second)
	t.Log(fmt.Sprintf("%02d", 1))
}
