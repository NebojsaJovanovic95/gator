-- name: CreateFeedFollow :one
with inserted_feed_follow as (

  INSERT INTO feed_follows (
    id, created_at, updated_at, user_id, feed_id
  ) VALUES ($1, $2, $3, $4, $5) RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users on users.id = inserted_feed_follow.user_id
INNER JOIN feeds on feeds.id = inserted_feed_follow.feed_id
;

-- name: GetFeedFollowsForUser :many
SELECT * FROM feed_follows WHERE user_id IN (SELECT id from users where name = $1);

