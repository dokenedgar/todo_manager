package main

import (
	"todo_manager/internal/todo"
)

func main() {
	todo.LoadTodo()
	defer todo.SaveChangesToFile()
	exit := false
	for !exit {
		input := todo.ReadFromCli(todo.InitialMenuPrompt(true))
		parsedInput := todo.ParseInput(input)
		for parsedInput > 5 || parsedInput <= 0 {
			input = todo.ReadFromCli(todo.InitialMenuPrompt(false))
			parsedInput = todo.ParseInput(input)
		}
		exit = todo.ProcessInput(parsedInput)
		todo.SaveChangesToFile()
	}
}
