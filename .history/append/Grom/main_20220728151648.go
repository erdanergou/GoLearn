package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
  )

  type person struct {
	id   int
	name string
	age  int
}



func main() {
	
}