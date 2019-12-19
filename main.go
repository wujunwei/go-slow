package main

import (
	"fmt"
	"goslow/rate"
	"math"
	"net/http"
	"strings"
	"time"
)

var l = rate.Create(5, 1, 5*time.Second)

type myHandle struct {
}

func (m myHandle) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	d, err := l.Acquire()
	result := strings.Builder{}
	result.WriteString(d.String())
	if err != nil {
		result.WriteString(err.Error())
	}
	_, _ = writer.Write([]byte("I have been waiting :" + result.String()))
}
func main() {
	//http.Handle("/acquire", myHandle{})
	//err := http.ListenAndServe("localhost:8080", nil)
	//fmt.Println(err)
	sum := 0.0
	account := 0.0
	var rater float64
	for i := 0; i < 36; i++ {
		if sum > 10000 {
			rater = 0.024
		} else {
			rater = 0.04
		}
		sum = sum * math.Pow(rater/365+1, 30)
		if sum > 10000 {
			account += 710
			if account > 1000 {
				sum += 1000
				account = account - 1000
			}
		} else {
			sum += 710
		}

		fmt.Printf("第%d个月收益为 %f元\n", i+1, sum)
	}
	fmt.Printf("还剩%f 元未加入计算利息", account)
}
