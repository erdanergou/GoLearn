package main

import "net/http"

// net/http server

func main() {
	http.HandleFunc("/posts/Go/15_socket/",)
	http.ListenAndServe("127.0.0.1:9090")
}
