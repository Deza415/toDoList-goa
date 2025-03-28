package todoapi

import (
	"context"

	todo "github.com/Deza415/toDoList-goa/gen/todo"
	"goa.design/clue/log"
)

// todo service example implementation.
// The example methods log the requests and return zero values.
type todosrvc struct{}

// NewTodo returns the todo service implementation.
func NewTodo() todo.Service {
	return &todosrvc{}
}

// List all todo items
func (s *todosrvc) List(ctx context.Context) (res []*todo.Todo, err error) {
	log.Printf(ctx, "todo.list")
	return
}

// Create a new todo item
func (s *todosrvc) Create(ctx context.Context, p *todo.CreatePayload) (res *todo.Todo, err error) {
	res = &todo.Todo{}
	log.Printf(ctx, "todo.create")
	return
}
