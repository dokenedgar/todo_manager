<!-- FIXME: write by following the Markdown syntax convention. Download an extension for `.md` files -->

# Todo Manager

This project is a CLI app for managing a list of todo items, stored in a file.

- [Todo Manager](#todo-manager)
  - [Installation](#installation)
  - [Introduction](#introduction)
  - [Features](#features)
    - [Add todo](#add-todo)
    - [Edit todo](#edit-todo)
    - [Delete todo](#delete-todo)
    - [View all todos](#view-all-todos)
    - [Exit](#exit)

## Installation

Clone the repo via

```bash
git clone https://github.com/dokenedgar/todo_manager.git
cd todo_manager
go run ./cmd/todo
```

## Introduction

Running the project, you get presented with 5 options (menu):

<!-- no toc -->
1. [Add todo](#add-todo)
2. [Edit todo](#edit-todo)
3. [Delete todo](#delete-todo)
4. [View all todos](#view-all-todos)
5. [Exit](#exit)
<!-- no toc -->

and also a prompt to respond with the number of the action to perform. This action (presenting the menu) is autmatically repeated and looped through, unti you choose to exit the program by selecting the 5 option, `Exit`

## Features

### Add todo

This allows you to add a todo to the existing ones, or create the first one if the list is empty.
You will be prompted to type the todo, and the result will be saved in a `json` file (todos.json), with the rest of the todos. The file is in the project's root directory.
The todos are saved in the order they're added, with the last one appended to the end of the list.
Each todo in the file is in this format

```json
{
  "title": "Todo description"
}
```

### Edit todo

This allows you to modify a todo at a particular index. When this option gets selected, you are further shown the list of current todos, and then prompted to select the number of the todo to modify. After which you can provide the updated todo, and it gets inserted at the specified index, not the end of the list.

### Delete todo

This will show the current todo items, with index numbers, and asks you to input which todo to delete, similar to when editing a todo.
The only difference is that here, the todo at the specified index gets deleted.

### View all todos

This will print all the todos present in the todos file, on the terminal.

### Exit

This terminates the program
