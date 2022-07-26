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
	// 对于get请求，
	fmt.Println(re.URL)
	fmt.Println(re.Method)
	fmt.Println(ioutil.ReadAll(re.Body)) // 我在服务端读出服务器端返回的响应的body
	rw.Write([]byte("hello"))

}

func main() {

	http.HandleFunc("/posts/Go/15_socket/", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
	// http.ListenAndServe("0.0.0.0:9090", nil) // 广播，全网可访问
}
