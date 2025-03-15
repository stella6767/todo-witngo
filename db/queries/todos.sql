-- db/queries/todos.sql
-- name: CreateTodo :one
INSERT INTO todos (user_id, title)
VALUES ($1, $2)
    RETURNING *;

-- name: GetTodosByUser :many
SELECT * FROM todos
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateTodoStatus :exec
UPDATE todos
SET completed = $2
WHERE id = $1;