-- name: CreateFeedFollow :many
WITH insert_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    ) RETURNING *
)
SELECT 
    insert_feed_follow.*,
    feeds.name AS feed_name
    users.name AS user_name
FROM feeds
INNER JOIN users ON users.id = $4,
INNER JOIN feeds ON feed.id = $5
