package main

import (
	"encoding/json"
	"log"
)

// mapstructure用于将通用的map[string]interface{}解码到对应的 Go 结构体中，或者执行相反的操作。

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}

type Cat struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	datas := []string{
		`{
			"type":"person",
			"name":"zhangsan",
			"age":"15",
			"job":"teacher"
		}`,
		`{
			"type":"cat",
			"name":"caidan",
			"age":"13"
		}`,
	}
	var person Person
	var cat Cat
	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}
	}
	
}
