package main

import (
	"firstapi/app"
	"firstapi/todo"
	"fmt"
)

func main() {
	todoList := todo.NewList()
	Handler := app.NewHandler(todoList)
	Server := app.NewServer(Handler)
	fmt.Println("START!!!")
	if err := Server.Start(); err != nil {
		fmt.Println(err)
	}
}
