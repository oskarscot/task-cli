package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/oskarscot/task-cli/internal/task"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	taskManager, err := task.LoadTasks()

	if err != nil {
		fmt.Printf("Error whilst loading the task manager: %v", err)
		return
	}

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		operation := scanner.Text()

		switch operation {
		case "add":
			addTask(taskManager, scanner)
		case "list":
			listTasks(taskManager)
		case "delete":
			deleteTask(taskManager, scanner)
		case "complete":
			completeTask(taskManager, scanner)
		case "save":
			saveTasks(taskManager)
		case "exit":
			return
		default:
			fmt.Printf("Unknown option: %s\n", operation)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Input error: %v\n", err)
	}
}

func saveTasks(t *task.TaskManager) {
	if err := t.SaveTasks(); err != nil {
		fmt.Printf("Failed to save tasks: %v\n", err)
		return
	}
	fmt.Println("Saved tasks!")
}

func completeTask(t *task.TaskManager, s *bufio.Scanner) {
	fmt.Print("Task ID: ")

	if s.Scan() {
		input := s.Text()
		taskId, err := strconv.Atoi(input)

		if err != nil {
			fmt.Printf("Input error: %v\n", err)
			return
		}

		if err := t.CompleteTask(taskId); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Task %d marked complete :DDD\n", taskId)
	}
}

func deleteTask(t *task.TaskManager, s *bufio.Scanner) {
	fmt.Print("Task ID: ")

	if s.Scan() {
		name := s.Text()
		taskId, err := strconv.Atoi(name)

		if err != nil {
			fmt.Printf("Input error: %v\n", err)
			return
		}

		if err := t.DeleteTask(taskId); err != nil {
			fmt.Printf("Error deleting task: %v\n", err)
			return
		}

		fmt.Printf("Deleted task %d\n", taskId)
	}
}

func listTasks(t *task.TaskManager) {
	fmt.Print(t.ListTasks())
}

func addTask(t *task.TaskManager, s *bufio.Scanner) {
	fmt.Print("Task Name: ")

	if s.Scan() {
		name := s.Text()
		t.AddTask(name)
		fmt.Printf("Added new task: %s\n", name)
	}
}
