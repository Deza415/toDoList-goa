package todoapi

import (
	"context"
	"fmt"

	todo "github.com/Deza415/toDoList-goa/gen/todo"
	"goa.design/clue/log"
)

// todo service example implementation.
// The example methods log the requests and return zero values.
type todosrvc struct{}

//NewTodo returns the todo service implementation.

func MyRealTodoService() todo.Service {
	return &todosrvc{}
}

// List all todo items
func (s *todosrvc) List(ctx context.Context) (res []*todo.Todo, err error) {
	log.Printf(ctx, "todo.list")
	return
}

// create
func (s *todosrvc) Create(ctx context.Context, p *todo.CreatePayload) (*todo.Todo, error) {
	fmt.Println("🔥 My real Create() was hit 🔥")
	fmt.Println("Received title:", p.Title)

	t := &todo.Todo{
		ID:        1,
		Title:     p.Title,
		Completed: false,
	}
	return t, nil
}
