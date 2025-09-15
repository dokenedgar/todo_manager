package main

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
