package main

type Person struct {
	Name string `gorm:"name" json:"name"`
	Age  int
	Id   int
}
