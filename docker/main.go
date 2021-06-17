package main

import (
	"goapp/app/src"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
)

func main() {
	src.StartServer()
}
