package todo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO: Why don't we load the file content in a slice on the program startup?
// Then, we can change stuff in the slice. Upon completion of the program, we can persist the change in the file.

var fileName = "todos.json"
var allTodos []TodoItem

func Add(todo string) {
	newTodo := TodoItem{
		Title: strings.ReplaceAll(todo, "\n", ""),
	}
	allTodos = append(allTodos, newTodo)
}

func ViewTodos() {
	if allTodos == nil {
		fmt.Println("Loading todos...")
		allTodos = getCurrentItems(fileName)
	}
	for i, todo := range allTodos {
		if i == 0 {
			fmt.Println("")
		}
		fmt.Printf("%d: %s\n", i+1, todo.Title)
	}
}

func ReadFromCli(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	// FIXME: always check for errors.
	line, err := reader.ReadString('\n')
	for err != nil {
		fmt.Printf("\nError occurred while reading input %s.\nPlease enter the value again", err.Error())
		fmt.Println(prompt)
		line, err = reader.ReadString('\n')
	}
	return line
}

func CreateFile(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0o644)
}

func checkFileExists(fileName string) bool {
	fileExists := false
	_, err := os.Stat(fileName)
	if err == nil {
		fileExists = true
	}
	return fileExists
}

func OpenFile(fileName string) (*os.File, error) {
	fileExists := checkFileExists(fileName)
	if fileExists {
		return os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0o644)
	}
	return CreateFile(fileName)
}

func WriteToFile(fileName string) error {
	file, errorToReturn := OpenFile(fileName)

	if file != nil {
		defer file.Close()
		data, err := json.MarshalIndent(allTodos, "", " ")
		if err != nil {
			return err
		}
		_, errorToReturn = file.Write(data)
		if errorToReturn != nil {
			return errorToReturn
		}
	}
	return errorToReturn
}

func getCurrentItems(fileName string) []TodoItem {
	var todos []TodoItem
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&todos); err != nil {
		if err.Error() != "EOF" {
			panic(err)
		}
	}
	return todos
}

func EditTodo() {
	ViewTodos()
	chosenTodo := ReadFromCli("\nEnter the number of the todo to edit:")
	fmt.Println("Editing number ", chosenTodo)
	strInt, err := strconv.Atoi(strings.ReplaceAll(chosenTodo, "\n", ""))
	for err != nil || strInt <= 0 || strInt-1 > len(allTodos)-1 {
		fmt.Println("\nPlease provide a valid number based on the number of todo you want to edit")
		chosenTodo := ReadFromCli("\nEnter the number of the todo to edit again:")
		strInt, err = strconv.Atoi(strings.ReplaceAll(chosenTodo, "\n", ""))
	}
	todoItem := allTodos[strInt-1]
	editedTodo := ReadFromCli("Enter the new todo in place of: " + todoItem.Title)
	editedTodoItem := TodoItem{
		Title: strings.ReplaceAll(editedTodo, "\n", ""),
	}
	var intermediate []TodoItem
	intermediate = allTodos[:strInt-1]
	intermediate = append(intermediate, editedTodoItem)
	intermediate = append(intermediate, allTodos[strInt:]...)
	allTodos = intermediate
}

func DeleteTodo() {
	ViewTodos()
	strIndexToDelete := ReadFromCli("\nEnter the number of the todo to delete (enter * to delete everything):")
	strIndexToDelete = strings.ReplaceAll(strIndexToDelete, "\n", "")
	if strIndexToDelete == "*" {
		fmt.Println("Deleting everthing!!")
	} else {
		strInt, err := strconv.Atoi(strIndexToDelete)
		for err != nil || strInt <= 0 || strInt-1 > len(allTodos)-1 {
			fmt.Println("\nPlease provide a valid number based on the number of todos")
			chosenTodo := ReadFromCli("\nEnter the number of the todo to delete again:")
			strInt, err = strconv.Atoi(strings.ReplaceAll(chosenTodo, "\n", ""))
		}
		todoItem := allTodos[strInt-1]
		var intermediate []TodoItem
		intermediate = allTodos[:strInt-1]
		intermediate = append(intermediate, allTodos[strInt:]...)
		allTodos = intermediate
		fmt.Println("Todo deleted successfully: ", todoItem)
	}
}

func LoadTodo() {
	allTodos = getCurrentItems(fileName)
}

func SaveChangesToFile() {
	fmt.Println("\nFinal save called")
	WriteToFile(fileName)
}
