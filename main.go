package main

import (
	"fmt"
	"github.com/wujunwei/go-slow/rate"
	"sync"
	"time"
)

func main() {
	l := rate.Create(5, 10*time.Second)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(l.Acquire())
			wg.Done()
		}()

	}
	wg.Wait()
}
