package main

import (
	"errors"
	"os"
	"testing"
)

func TestToggle(t *testing.T) {
	todos := Done{}
	todos.add("Test task")

	if err := todos.toggle(0); err != nil {
		t.Fatalf("toggle failed: %v", err)
	}

	if !todos[0].Completed {
		t.Fatal("expected todo to be completed")
	}

	if todos[0].CompletedAt == nil {
		t.Fatal("expected CompletedAt to be set")
	}

	if err := todos.toggle(0); err != nil {
		t.Fatalf("toggle failed: %v", err)
	}

	if todos[0].Completed {
		t.Fatal("expected todo to be incomplete")
	}

	if todos[0].CompletedAt != nil {
		t.Fatal("expected CompletedAt to be nil")
	}
}

func TestEdit(t *testing.T) {
	todos := Done{}
	todos.add("Old title")

	if err := todos.edit(0, "New title"); err != nil {
		t.Fatalf("edit failed: %v", err)
	}

	if todos[0].Title != "New title" {
		t.Fatalf("expected title %q, got %q", "New title", todos[0].Title)
	}

	if err := todos.edit(999, "Invalid"); err == nil {
		t.Fatal("expected error for invalid index")
	}
}

func TestDelete(t *testing.T) {
	todos := Done{}
	todos.add("Task 1")
	todos.add("Task 2")

	if err := todos.delete(0); err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	if len(todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(todos))
	}

	if todos[0].Title != "Task 2" {
		t.Fatalf("expected remaining task to be %q, got %q", "Task 2", todos[0].Title)
	}

	if err := todos.delete(999); err == nil {
		t.Fatal("expected error for invalid index")
	}
}

func TestAdd(t *testing.T) {
	todos := Done{}

	todos.add("New task")

	if len(todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(todos))
	}

	if todos[0].Title != "New task" {
		t.Fatalf("expected title %q, got %q", "New task", todos[0].Title)
	}

	if todos[0].Completed {
		t.Fatal("expected new todo to be incomplete")
	}
	if todos[0].CompletedAt != nil {
		t.Fatal("expected CompletedAt to be nil")
	}

	if todos[0].CreatedAt.IsZero() {
		t.Fatal("expected CreatedAt to be set")
	}
}

func TestValidateIndex(t *testing.T) {
	todos := Done{}
	todos.add("Task 1")

	validIndexes := []int{0}
	for _, index := range validIndexes {
		if err := todos.validateIndex(index); err != nil {
			t.Fatalf("expected index %d to be valid, got error: %v", index, err)
		}
	}

	invalidIndexes := []int{-1, 1, 999}
	for _, index := range invalidIndexes {
		if err := todos.validateIndex(index); !errors.Is(err, ErrInvalidIndex) {
			t.Fatalf("expected ErrInvalidIndex for index %d, got %v", index, err)
		}
	}
}

func TestStorage(t *testing.T) {
	file := t.TempDir() + "/todos.json"

	todos := Done{}
	todos.add("Storage test")

	storage := NewStorage[Done](file)

	if err := storage.Save(todos); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	loaded := Done{}

	if err := storage.Load(&loaded); err != nil {
		t.Fatalf("load failed: %v", err)
	}

	if len(loaded) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(loaded))
	}

	if loaded[0].Title != "Storage test" {
		t.Fatalf("expected title %q, got %q", "Storage test", loaded[0].Title)
	}
}

func TestStorageLoadMissingFile(t *testing.T) {
	file := t.TempDir() + "/missing.json"
	storage := NewStorage[Done](file)
	todos := Done{}

	err := storage.Load(&todos)

	if err == nil {
		t.Fatal("expected error when loading missing file")
	}

	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected os.ErrNotExist, got %v", err)
	}
}

func TestClear(t *testing.T) {
	todos := Done{}

	todos.add("Task 1")
	todos.add("Task 2")
	todos.add("Task 3")

	if len(todos) != 3 {
		t.Fatalf("expected 3 todos, got %d", len(todos))
	}

	todos.clear()

	if len(todos) != 0 {
		t.Fatalf("expected 0 todos after clear, got %d", len(todos))
	}
}
