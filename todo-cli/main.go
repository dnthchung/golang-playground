package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	ID   int
	Task string
	Done bool
}

var todos []Todo
var nextID int = 1

func main() {
	fmt.Println("Welcome to the Todo CLI App!")
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add a task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Mark a task as done")
		fmt.Println("4. Delete a task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
			continue
		}

		switch choice {
		case 1:
			addTask()
		case 2:
			listTasks()
		case 3:
			markTaskAsDone()
		case 4:
			deleteTask()
		case 5:
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}

func addTask() {
	fmt.Print("Enter the task description: ")
	reader := bufio.NewReader(os.Stdin)
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)
	if task == "" {
		fmt.Println("Task description cannot be empty.")
		return
	}
	todos = append(todos, Todo{ID: nextID, Task: task, Done: false})
	fmt.Printf("Task added with ID: %d\n", nextID)
	nextID++
}

func listTasks() {
	if len(todos) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("\nTodo List:")
	for _, todo := range todos {
		status := "[ ]"
		if todo.Done {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", todo.ID, status, todo.Task)
	}
}

func markTaskAsDone() {
	fmt.Print("Enter the task ID to mark as done: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	taskID, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid task ID.")
		return
	}

	for i, todo := range todos {
		if todo.ID == taskID {
			if todos[i].Done {
				fmt.Println("Task is already marked as done.")
			} else {
				todos[i].Done = true
				fmt.Println("Task marked as done.")
			}
			return
		}
	}
	fmt.Println("Task not found.")
}

func deleteTask() {
	fmt.Print("Enter the task ID to delete: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	taskID, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid task ID.")
		return
	}

	for i, todo := range todos {
		if todo.ID == taskID {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("Task deleted.")
			return
		}
	}
	fmt.Println("Task not found.")
}
