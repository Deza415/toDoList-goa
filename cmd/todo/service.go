package main

import (
	"context"
	"fmt"

	"github.com/Deza415/toDoList-goa/gen/todo"
)

// In-memory todo storage
var todos []*todo.Todo
var nextID int32 = 1

// todosrvc implements the todo.Service interface
type todosrvc struct{}

// NewTodo returns a new todo service instance
func NewTodo() todo.Service {
	return &todosrvc{}
}

// Create adds a new todo
func (s *todosrvc) Create(ctx context.Context, p *todo.CreatePayload) (*todo.Todo, error) {
	fmt.Println("🔥 My real Create() was hit 🔥")
	fmt.Println("Received title:", p.Title)

	t := &todo.Todo{
		ID:        nextID,
		Title:     p.Title,
		Completed: false,
	}
	nextID++
	todos = append(todos, t)
	return t, nil
}

// List returns all todos
func (s *todosrvc) List(ctx context.Context) ([]*todo.Todo, error) {
	return todos, nil
}
