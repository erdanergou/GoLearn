package main

// go语言中内置的map并不是并发安全的
var m = make(map[string]int)

func get(key string)int{
	return m[key]
}

func set(key string,value int){
	m[key] = value
}

func main() {
	wg := 
}
