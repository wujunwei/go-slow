package main

import (
	"fmt"
	"goslow/rate"
	"net/http"
	"strings"
	"time"
)

var l = rate.Create(5, 0, 5*time.Second)

type myHandle struct {
}

func (m myHandle) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	d, err := l.Acquire()
	result := strings.Builder{}
	if err == nil {
		result.WriteString(d.String())
	} else {
		result.WriteString(err.Error())
	}
	_, _ = writer.Write([]byte("I have been waiting :" + result.String()))
}
func main() {
	//http.Handle("/", myHandle{})
	//err := http.ListenAndServe("localhost:8080", nil)
	//fmt.Println(err)
	for i := 0; i < 10; i++ {
		go func() {
			t, _ := l.Acquire()
			fmt.Println("I have been waiting :" + t.String())
		}()

	}
	time.Sleep(7 * time.Second)
}
