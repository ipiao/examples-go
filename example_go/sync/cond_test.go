package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})

	count := int64(5)

	// 生产者
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			cond.L.Lock()
			if count >= 10 {
				fmt.Println("channel is full!!")
				//cond.Wait()
			} else {
				count++
				if count == 1 {
					cond.Signal()
				}
				fmt.Printf("produce one, now is %d\n!!", count)
			}
			cond.L.Unlock()
		}
	}()

	// 消费者
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 10)
		go func(n int) {
			for {
				time.Sleep(time.Millisecond * 500)
				cond.L.Lock()
				if count == 0 {
					fmt.Printf("consumer %d find channel empty!!", n)
					cond.Wait()
				} else {
					fmt.Printf("consumer %d take one, now is %d\n!!", n, count)
					count--
				}
				cond.L.Unlock()
			}
		}(i)
	}

	select {}
}
