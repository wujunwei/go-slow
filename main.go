package main

import (
	"fmt"
	"github.com/wujunwei/go-slow/rate"
	"time"
)

func main() {
	l := rate.Create(10, time.Second)
	for i := 0; i < 1000; i++ {
		fmt.Println(l.Acquire())
	}

}
