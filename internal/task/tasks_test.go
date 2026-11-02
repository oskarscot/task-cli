package task

import "testing"

func TestAddTask(t *testing.T) {
	manager := NewTaskManager()

	manager.AddTask("Learn Go")

	if len(manager.tasks) != 1 {
		t.Fatalf(
			"expected 1 task, got %d",
			len(manager.tasks),
		)
	}

	if manager.tasks[0].Name != "Learn Go" {
		t.Errorf(
			"expected task name Learn Go, got %s",
			manager.tasks[0].Name,
		)
	}
}

func TestDeleteTask(t *testing.T) {
	manager := NewTaskManager()

	manager.AddTask("hello")

	if err := manager.DeleteTask(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := len(manager.tasks); got != 0 {
		t.Errorf("expected length 0, got %d", got)
	}
}
