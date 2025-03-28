package design

import . "goa.design/goa/v3/dsl"

var _ = Service("todo", func() {
	Description("The todo service manages todo items")

	Method("list", func() {
		Description("List all todo items")
		Result(ArrayOf(Todo))
		HTTP(func() {
			GET("/")
		})
	})

	Method("create", func() {
		Description("Create a new todo item")
		Payload(CreatePayload)
		Result(Todo)
		HTTP(func() {
			POST("/")
			Body(func() {
				Attribute("title")
			})
		})
	})
})

var Todo = Type("Todo", func() {
	Description("A single todo item")
	Field(1, "id", Int32, "Unique ID")
	Field(2, "title", String, "Title of the todo")
	Field(3, "completed", Boolean, "Whether the todo is completed")
	Required("id", "title", "completed")
})

var CreatePayload = Type("CreatePayload", func() {
	Description("Payload for creating a todo item")
	Field(1, "title", String, "Title of the todo")
	Required("title")
})
