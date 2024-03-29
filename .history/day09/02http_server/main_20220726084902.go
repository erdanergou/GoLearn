package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

func f1(rw http.ResponseWriter, re *http.Request) {
	str, err := ioutil.ReadFile("./index.html")
	if err != nil {
		rw.Write([]byte(err.Error()))
	}
	rw.Write(str)
}

func f2(rw http.ResponseWriter, re *http.Request) {
	fmt.Println(re.URL)
	fmt.Println(re.Method)
	fmt.Println(ioutil.ReadAll(re.Body))
	rw.Write([]byte("hello"))

}

func main() {

	http.HandleFunc("/posts/Go/15_socket/", f1)
	http.HandleFunc("/posts/Go/15_socket/", f1)
	http.ListenAndServe("127.0.0.1:9090", nil)
	// http.ListenAndServe("0.0.0.0:9090", nil)  // 广播，全网可访问
}
