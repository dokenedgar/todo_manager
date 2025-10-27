package todo

import (
	"fmt"
	"strconv"
	"strings"
)

type TodoItem struct {
	Title string `json:"title"`
}

// FIXME: all the exported functions must have a comment. See warning ST1000.
// FIXME: this file can be renamed as "manager.go". "todo_manager" is redundant with the pkg name (and folder name).

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
	// BUG: this default is useless since if we're here it means we have an input ranging from 1 to 5.
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
