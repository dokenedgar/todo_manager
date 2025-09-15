package main
import "todo_manager/internal/todo"

func main() {
	input := todo.ReadFromCli(`
	What do you want to do? Type only a number: 

	1 Add todo
	2 Edit todo
	3 Delete todo
	4 View all todos
	
	`)

	todo.ProcessInput(input)
}
