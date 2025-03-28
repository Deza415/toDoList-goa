// Code generated by goa v3.20.1, DO NOT EDIT.
//
// todo endpoints
//
// Command:
// $ goa gen github.com/Deza415/toDoList-goa/design

package todo

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "todo" service endpoints.
type Endpoints struct {
	List   goa.Endpoint
	Create goa.Endpoint
}

// NewEndpoints wraps the methods of the "todo" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List:   NewListEndpoint(s),
		Create: NewCreateEndpoint(s),
	}
}

// Use applies the given middleware to all the "todo" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Create = m(e.Create)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "todo".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.List(ctx)
	}
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "todo".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreatePayload)
		return s.Create(ctx, p)
	}
}
