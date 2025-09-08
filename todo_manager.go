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
		fmt.Println("Add todo")
		todo := ReadFromCli("Type in your todo to add")
		Add(fileName, todo)
		fmt.Println("\nTodo added successfully")
	case 2:
		fmt.Println("\n***** Edit todo *****")
		ViewTodos()
		chosenTodo := ReadFromCli("\nEnter the number of the todo to edit:")
		fmt.Println("Editing number ", chosenTodo)
		EditTodo(chosenTodo)
	case 3:
		fmt.Println("Delete todo")
	case 4:
		fmt.Println("\n***** View all todos *****")
		ViewTodos()
	default:
		fmt.Println("I dunno what you wanna do hommie")
	}
}
