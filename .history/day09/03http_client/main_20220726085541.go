package main

//http client

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("127.0.0.1:9090/xxx/")
	if err != nil {
		fmt.Println("get url failed,err:", err)
		return
	}
	resp.

}


