package main

import (
	"todo_manager/internal/todo"
)

// FIXME: when I type "5" to exit, it will display "Final save called" two times.
// FIXME: 'todos.json' should not be added to Git.

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
