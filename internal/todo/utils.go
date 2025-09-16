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
	line, _ := reader.ReadString('\n')
	return line
}

func CreateFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}
	return file, nil
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
	var file *os.File
	var err error
	if fileExists {
		file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0o644)
	} else {
		file, err = CreateFile(fileName)
	}
	return file, err
}

func WriteToFile(fileName string) error {
	file, errorToReturn := OpenFile(fileName)

	if file != nil {
		defer file.Close()
		data, err := json.MarshalIndent(allTodos, "", " ")
		if err != nil {
			return err
		} else {
			_, errorToReturn = file.Write(data)
			if errorToReturn != nil {
				return errorToReturn
			}
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

func EditTodo(indexToEdit string) {
	strInt, err := strconv.Atoi(strings.ReplaceAll(indexToEdit, "\n", ""))
	if err != nil {
		panic(err)
	}
	todos := getCurrentItems(fileName)
	if strInt-1 > len(todos)-1 {
		panic("Chosen number greater than number of items in the list")
	}
	todoItem := todos[strInt-1]
	editedTodo := ReadFromCli("Enter the new todo in place of: " + todoItem.Title)
	editedTodoItem := TodoItem{
		Title: strings.ReplaceAll(editedTodo, "\n", ""),
	}
	var intermediate []TodoItem
	intermediate = todos[:strInt-1]
	intermediate = append(intermediate, editedTodoItem)
	intermediate = append(intermediate, todos[strInt:]...)
	fmt.Println(intermediate)
	err = WriteToFile(fileName)
	if err != nil {
		panic(err)
	}
}

func DeleteTodo(strIndexToDelete string) {
	strIndexToDelete = strings.ReplaceAll(strIndexToDelete, "\n", "")
	if strIndexToDelete == "*" {
		fmt.Println("Deleting everthing!!")
	} else {
		strInt, err := strconv.Atoi(strIndexToDelete)
		if err != nil {
			panic(err)
		}
		todos := getCurrentItems(fileName)
		if strInt-1 > len(todos)-1 {
			panic("Chosen number greater than number of items in the list")
		}
		todoItem := todos[strInt-1]

		var intermediate []TodoItem
		intermediate = todos[:strInt-1]
		intermediate = append(intermediate, todos[strInt:]...)
		fmt.Println(intermediate)
		err = WriteToFile(fileName)
		if err != nil {
			panic(err)
		}
		fmt.Println("Todo deleted successfully: ", todoItem)
	}
}

func LoadTodo() {
	allTodos = getCurrentItems(fileName)
}

func SaveChangesToFile() {
	fmt.Println("Final save called")
	WriteToFile(fileName)
}
