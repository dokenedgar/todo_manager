package todo

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
var allTodos []TodoItem

// TODO: I want to be able to perform more than an operation. Provide me with the "Exit" option for when I want to exit.

func ProcessInput(input int) {
	switch input {
	case 1:
		fmt.Println("\n***** Add todo *****")
		todo := ReadFromCli("Type in your todo to add:\n")
		Add(todo)
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
	}
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
	
	`
	if !initialPrompt {
		promptString = `
	Invalid input, please provide a value within the following options shown.

	What do you want to do? Type only a number: 

	1 Add todo
	2 Edit todo
	3 Delete todo
	4 View all todos
	
	`
	}
	return promptString
}
