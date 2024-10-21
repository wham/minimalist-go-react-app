package storage

import "testing"

func TestTodos(t *testing.T) {
	todos := Todos{}

	if todos.List() != nil {
		t.Fatalf("expected nil")
	}

	todos.Add(Todo{"Clean the dishes"})
	todos.Add(Todo{"Read a book"})

	if len(todos.List()) != 2 {
		t.Fatalf("exp 2")
	}
}
