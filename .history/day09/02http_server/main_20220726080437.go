package main

import "net/http"

// net/http server

func main() {
	http.HandleFunc("/posts/Go/15_socket/")
	http.ListenAndServe()
}
