package main

import (
	"io/ioutil"
	"net/http"
)

// net/http server

func f1(rw http.ResponseWriter, re *http.Request) {
	ioutil.ReadFile("")
	rw.Write([]byte(str))
}

func main() {

	http.HandleFunc("/posts/Go/15_socket/", f1)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
