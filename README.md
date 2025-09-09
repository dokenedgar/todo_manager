# Todo Manager
This project allows managing a list of todo items on a file, via the terminal. 
It is a CLI app that can be ran with `go run .`

## Introduction
Running the project, a user gets presented with 4 options (menu):

- 1 Add todo
-	2 Edit todo
-	3 Delete todo
-	4 View all todos

and also a prompt to respond with the number of the action to perform.

### 1. Add todo
This allows a user to add a todo to the existing ones, or create the first one if the list is empty.
The user will be prompted to type the todo, and the result will be saved in a `json` file (todos.json), with the rest of the todos.
The todos are saved in the order they're added, with the last one appended to the end of the list.
Each todo is in this format
```
  {
    "title": "Todo description"
  }
```

### 2. Edit todo
This allows you to modify a todo at a particular index. When this option gets selected, the user is further
shown the list of current todos, and then prompted to select the number of the todo to modify. After which 
the user can provide the updated toto, and it gets inserted at the specified index, not the end of the list.

### 3. Delete todo
The will show the current todo items, with an index number, and ask the user to input which todo to delete, similar to when editing a todo.
The only difference is here, the todo at the specified index gets deleted.

### 4. View all todos
This will print all the todos present in the todos file, on the terminal.
