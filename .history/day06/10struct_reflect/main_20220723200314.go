package main

//结构体反射
type student struct{
	Name string `json:"name"`
	Age int `json:"age"`
}


func main() {
	stu1 := student{
		Name:"小王子",
		Age:12
	}	
	t:=reflect.Typeof(stu1)
	fmt.Println(t.Name())
}