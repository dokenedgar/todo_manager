package todo

import (
	"fmt"
	"strconv"
	"strings"
)

type TodoItem struct {
	Title string `json:"title"`
}



// FIXME: the func 'main' of the 'main' pkg should stay in a main.go file
// FIXME: address all the warnings the staticcheck, or go vet is giving to you.
// FIXME: if I type "5" I want to stay in the loop. Until I type a valid input.
// Every time the user failed to input, share a message with the valid inputs.
// TODO: I want to be able to perform more than an operation. Provide me with the "Exit" option for when I want to exit.
// TODO: the files should be organized in a better way. Use the `cmd` and `internal` folders.

func ProcessInput(input int) bool {
	exit := false
	switch input {
	case 1:
		fmt.Println("\n***** Add todo *****")
		todo := ReadFromCli("Type in your todo to add:\n")
		Add(todo)
		fmt.Println("\nTodo added successfully")
	case 2:
		fmt.Println("\n***** Edit todo *****")
		EditTodo()
	case 3:
		fmt.Println("D\n***** Delete todo *****")
		DeleteTodo()
	case 4:
		fmt.Println("\n***** View all todos *****")
		ViewTodos()
	case 5:
		fmt.Println("\n***** Goodbye! *****")
		exit = true
	default: // This may not be needed, as it won't ever be called based on the current code in main.go
		// FIXME: '\n' at the end is redundant
		fmt.Println("\nI dunno what you wanna do hommie\nBut here're all the current todos:")
		ViewTodos()
	}
	return exit
}

func ParseInput(input string) int {
	intValue, err := strconv.Atoi(strings.ReplaceAll(input, "\n", ""))
	if err != nil {
		intValue = -1
	}
	return intValue
}

func InitialMenuPrompt(initialPrompt bool) string {
	promptString := `
	What do you want to do? Type only a number: 

	1 Add todo
	2 Edit todo
	3 Delete todo
	4 View all todos
	5 Exit
	
	`
	if !initialPrompt {
		promptString = `
	Invalid input, please provide a value within the following options shown.

	What do you want to do? Type only a number: 

	1 Add todo
	2 Edit todo
	3 Delete todo
	4 View all todos
	5 Exit
	
	`
	}
	return promptString
}
