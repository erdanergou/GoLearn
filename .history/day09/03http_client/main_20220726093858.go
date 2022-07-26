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

	//自定义请求
	data := url.Values{} // url encode
	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name", "小垃圾")
	data.Set("age", "12")
	queryStr := data.Encode() // url encode 之后的url
	fmt.Println(queryStr)
	urlObj.Query = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("get url failed ,err: ", err)
		return
	}
	
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	client.Do(req)
	defer resp.Body.Close() //必须关闭resp.Body
	// 发请求

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
