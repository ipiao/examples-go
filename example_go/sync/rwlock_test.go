package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/panjf2000/ants"
)

// 测试读写锁的饥饿等待问题
func TestRWLock(t *testing.T) {
	mu := sync.RWMutex{}

	i := 0

	read := func(interface{}) {
		i++
		fmt.Printf("read %d\n", i)
		time.Sleep(time.Millisecond * 10) // 每个read执行10ms
	}

	pw, err := ants.NewPoolWithFunc(1000000, read)
	if err != nil {
		t.Fatal(err)
	}
	done := make(chan struct{})
	tick := time.NewTicker(time.Millisecond)

	go func() {
		time.Sleep(time.Second) // 1s 之后执行write
		mu.Lock()
		defer mu.Unlock()
		fmt.Println("this is a write lock")
		close(done)
		return
	}()

	for {
		select {
		case <-tick.C:
			pw.Invoke(struct{}{})
		case <-done:
			return
		}
	}
}
