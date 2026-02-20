-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

/* 
Values of $1-$6 correspond to four different parameters that can be passed into the query in Go code
the :one on the line 1 are telling SQLC to expect a single row returned because of the RETURNING * at the end
*/