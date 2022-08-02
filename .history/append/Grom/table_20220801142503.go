package main

type Person struct {
	Name string `gorm:"name" json:"name"`
	Age  int	`gorm:"age" json:"age"`
	Id   int	``
}
