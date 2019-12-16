package main

import (
	"fmt"
	"github.com/wujunwei/go-slow/rate"
	"sync"
	"time"
)

func main() {
	l := rate.Create(5, 1, 1*time.Second)
	wg := sync.WaitGroup{}
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i, l.Acquire())
			wg.Done()
		}(i)

	}
	wg.Wait()
}
