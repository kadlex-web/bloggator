-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

/* 
Values of $1-$4 correspond to four different parameters that can be passed into the query in Go code
the :one on the line 1 are telling SQLC to expect a single row returned because of the RETURNING * at the end
*/