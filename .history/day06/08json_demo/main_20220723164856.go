package main

import "encoding/json"

// Json

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"张三","age":15}`
	var p person
	json.Unmarshal([]byte(str))
}
