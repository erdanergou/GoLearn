package main

type person struct {
	id   int
	name string
	age  int
}

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
  )

func main() {

}
