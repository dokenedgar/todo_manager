package main

import (
	"fmt"
	"strconv"
	"strings"
)

type DataToWrite struct {
	SingleItem   TodoItem   `json:"single_item"`
	ArrayOfItems []TodoItem `json:"array_items"`
}
type TodoItem struct {
	Title string `json:"title"`
}

var fileName = "todos.json"

// FIXME: the func 'main' of the 'main' pkg should stay in a main.go file
// FIXME: address all the warnings the staticcheck, or go vet is giving to you.
// FIXME: if I type "5" I want to stay in the loop. Until I type a valid input.
// Every time the user failed to input, share a message with the valid inputs.
// TODO: I want to be able to perform more than an operation. Provide me with the "Exit" option for when I want to exit.
// TODO: the files should be organized in a better way. Use the `cmd` and `internal` folders.
func main() {
	input := ReadFromCli(`
	What do you want to do? Type only a number: 

	1 Add todo
	2 Edit todo
	3 Delete todo
	4 View all todos
	
	`)

	processInput(input)
}

func processInput(input string) {
	intValue, err := strconv.Atoi(strings.ReplaceAll(input, "\n", ""))
	if err != nil {
		fmt.Println("Error occurred while parsing input: ", err)
		return
	}
	switch intValue {
	case 1:
		fmt.Println("\n***** Add todo *****")
		todo := ReadFromCli("Type in your todo to add:\n")
		Add(fileName, todo)
		fmt.Println("\nTodo added successfully")
	case 2:
		fmt.Println("\n***** Edit todo *****")
		ViewTodos()
		chosenTodo := ReadFromCli("\nEnter the number of the todo to edit:")
		fmt.Println("Editing number ", chosenTodo)
		EditTodo(chosenTodo)
	case 3:
		fmt.Println("D\n***** Delete todo *****")
		ViewTodos()
		chosenTodo := ReadFromCli("\nEnter the number of the todo to delete (enter * to delete everything):")
		DeleteTodo(chosenTodo)
	case 4:
		fmt.Println("\n***** View all todos *****")
		ViewTodos()
	default:
		// FIXME: '\n' at the end is redundant
		fmt.Println("\nI dunno what you wanna do hommie\nBut here're all the current todos:\n")
		ViewTodos()
	}
}
