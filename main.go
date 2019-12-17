package main

import (
	"fmt"
	"github.com/wujunwei/go-slow/rate"
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
	http.Handle("/acquire", myHandle{})
	err := http.ListenAndServe("localhost:8080", nil)
	fmt.Println(err)
}
