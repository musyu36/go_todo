package main

import (
	"fmt"
	"golang/todo_app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// init 関数を呼ぶため
	fmt.Println(models.Db)
}
