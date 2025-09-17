package main

import (
	"todo_manager/internal/todo"
)

func main() {
	todo.LoadTodo()
	defer todo.SaveChangesToFile()
	input := todo.ReadFromCli(todo.InitialMenuPrompt(true))
	parsedInput := todo.ParseInput(input)
	for parsedInput > 4 || parsedInput <= 0 {
		input = todo.ReadFromCli(todo.InitialMenuPrompt(false))
		parsedInput = todo.ParseInput(input)
	}

	todo.ProcessInput(parsedInput)
}
