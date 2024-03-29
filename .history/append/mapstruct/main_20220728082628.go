package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

// mapstructure用于将通用的map[string]interface{}解码到对应的 Go 结构体中，或者执行相反的操作。

type Person struct {
	Name string `mapstructure:"username"`
	Age  int
	Job  string
}

type Cat struct {
	Name  string
	Age   int
	Breed string
}

func main() {
	datas := []string{`
	  { 
		"type": "person",
		"username":"dj",
		"age":18,
		"job": "programmer"
	  }
	`,
		`
	  {
		"type": "cat",
		"name": "kitty",
		"Age": 1,
		"breed": "Ragdoll"
	  }
	`,
		`
	  {
		"type": "cat",
		"Name": "rooooose",
		"age": 2,
		"breed": "shorthair"
	  }
	`,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v\n", m)
		fmt.Printf("m[‘] %v\n", m["type"].(string))
		switch m["type"].(string) {
		case "person":
			var p Person
			mapstructure.Decode(m, &p)
			fmt.Println("person", p)

		case "cat":
			var cat Cat
			mapstructure.Decode(m, &cat)
			fmt.Println("cat", cat)
		}
	}
}
