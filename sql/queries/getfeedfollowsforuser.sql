-- name: GetFeedFollowsForUser :many
WITH userFollowing AS (
	SELECT * FROM feed_follows
	WHERE feed_follows.user_id = $1
)
SELECT
	userFollowing.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM userFollowing
INNER JOIN users ON users.id = userFollowing.user_id
INNER JOIN feeds ON feeds.id = userFollowing.feed_id;