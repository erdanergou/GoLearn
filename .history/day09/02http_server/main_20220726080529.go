package main

import "net/http"

// net/http server

func f1(http.ResponseWriter, http.Request) {

}

func main() {

	http.HandleFunc("/posts/Go/15_socket/")
	http.ListenAndServe("127.0.0.1:9090", nil)
}