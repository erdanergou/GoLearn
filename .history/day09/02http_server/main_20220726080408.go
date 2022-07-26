package main

import "net/http"

// net/http server

func main(){
	http.HandleFunc("")
	http.ListenAndServe()
}