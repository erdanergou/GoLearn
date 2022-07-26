package main

import "net/http"

// net/http server

func f1(http.ResponseWriter,*http.Request){
	str := "hello 沙河！"
	w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/posts/Go/15_socket/",f1)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
