# Todo REST API in Go

A simple REST API for managing todo items, built with Go and the [Gin](https://gin-gonic.com/) web framework. Todos are stored in memory — data resets when the server restarts.

---

## Project structure

| File | Description |
|---|---|
| `main.go` | Entry point. Initializes the Gin router and registers all 5 routes. |
| `handlers.go` | All route handler functions, the `todo` struct, the in-memory slice, and the ID counter. |

---

## Running the server

```bash
go mod tidy
go run .
```

The server starts at `http://localhost:8080`.

---

## Endpoints

| Method | Path | Description |
|---|---|---|
| GET | `/todos` | Return all todos |
| GET | `/todos/:id` | Return a single todo by ID |
| POST | `/todos` | Create a new todo |
| PUT | `/todos/:id` | Update a todo's title and/or completion status |
| DELETE | `/todos/:id` | Delete a todo |

---

## Testing with cURL

### Create a todo
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Finish Go project", "completed": false}'
```

### Get all todos
```bash
curl http://localhost:8080/todos
```

### Get a todo by ID
```bash
curl http://localhost:8080/todos/1
```

### Update a todo (One or both fields can be sent)
```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Updated title", "completed": true}'
```

### Delete a todo
```bash
curl -X DELETE http://localhost:8080/todos/1
```

---

## Error responses

| Status | Cause |
|---|---|
| `400` | Malformed JSON, invalid ID, or empty title |
| `404` | No todo found with the given ID |
