package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Task struct {
	ID       int
	Name     string
	Complete bool
}

type TaskManager struct {
	tasks  []Task
	nextId int
}

type TaskFile struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (t *TaskManager) AddTask(name string) {
	t.nextId++
	t.tasks = append(t.tasks, Task{
		ID:       t.nextId,
		Name:     name,
		Complete: false,
	})
}

func (t *TaskManager) ListTasks() string {
	var builder strings.Builder

	for _, task := range t.tasks {
		status := " "
		if task.Complete {
			status = "x"
		}

		fmt.Fprintf(&builder, "[%s] %d - %s\n", status, task.ID, task.Name)
	}

	return builder.String()
}

func (t *TaskManager) DeleteTask(id int) error {
	for index, task := range t.tasks {
		if task.ID == id {
			t.tasks = slices.Delete(t.tasks, index, index+1)
			return nil
		}
	}

	return fmt.Errorf("Task %d does not exist", id)
}

func (t *TaskManager) CompleteTask(id int) error {
	for index := range t.tasks {
		if t.tasks[index].ID == id {
			t.tasks[index].Complete = true
			return nil
		}
	}

	return fmt.Errorf("task %d does not exist", id)
}

func (t *TaskManager) SaveTasks() error {
	data, err := json.MarshalIndent(t.ToTaskFile(), "", "  ")

	if err != nil {
		return fmt.Errorf("failed to encode tasks: %w", err)
	}

	if err := os.WriteFile("data.json", data, 0o644); err != nil {
		return fmt.Errorf("failed to save tasks: %w", err)
	}

	return nil
}

func (t *TaskManager) ToTaskFile() *TaskFile {
	return &TaskFile{
		Tasks:  t.tasks,
		NextID: t.nextId,
	}
}

func LoadTasks() (*TaskManager, error) {
	var taskFile TaskFile
	bytes, err := os.ReadFile("data.json")

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("No saved task data found, resuming with fresh task manager!")
		return NewTaskManager(), nil
	}

	if err != nil {
		return nil, fmt.Errorf("Error whilst loading task data %w", err)
	}

	if err := json.Unmarshal(bytes, &taskFile); err != nil {
		return nil, fmt.Errorf("Error whilst decoding json: %w", err)
	}

	return &TaskManager{
		tasks:  taskFile.Tasks,
		nextId: taskFile.NextID,
	}, nil
}
