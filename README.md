# TodoList-Goa Backend

## The stuff so far

- Ran `goa init todolist` to "scaffold" the project
- Created `design/todo.go` to define the API using Goa DSL
- Ran `goa gen` to generate service definitions and transport layers
- Ran `goa example` to "scaffold" a working HTTP server
- Ran `go mod tidy` to resolve dependencies

- Ran the service with `go run ./cmd/todo` — it's live at http://localhost:80

- Added `Create` and `List` methods to the Goa DSL in `design/todo.go`
- Regenerated code with `goa gen` after DSL changes
- Created a real service implementation using `todosrvc`
- Defined `MyRealTodoService()` to return your custom implementation
- Updated `main.go` to call `todoapi.MyRealTodoService()` instead of any stubs
- Verified everything by printing from inside `Create()` using `fmt.Println`
- Used `curl` to send a POST request and confirmed that your logic runs
- Removed any unused stub implementations and cleaned up project structure
- Confirmed Goa is no longer using generated stubs by replacing with your own
