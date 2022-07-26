package main

//http client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// resp, err := http.Get("http://127.0.0.1:9090/xxx/?username=admin&password=1234")
	// if err != nil {
	// 	fmt.Println("get url failed,err:", err)
	// 	return
	// }

	data := url.Values{} // url encode
	data.Set("name", "小垃圾")
	data.Set("age", "12")
	data.
	http.NewRequest("GET", data)
	// 从resp中读取服务端返回的数据
	// resp.Body.Read()
	// resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body) // 我在客户端读出服务器端返回的响应的body
	if err != nil {
		fmt.Println("read resp.Body failed,err: ", err)
		return
	}
	fmt.Println(string(b))
}
