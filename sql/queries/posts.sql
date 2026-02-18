-- name: CreatePost :one
INSERT INTO posts (
  id, created_at, updated_at, title, url, description, published_at, feed_id
) VALUES ($1, $2, $3, $4, $5, $6, $7,
  (SELECT id from feeds where feeds.url = $5 LIMIT 1)
) RETURNING *;

-- name: GetPosts :many
SELECT * FROM posts where feed_id IN (
  select feed_id from feed_follows where user_id = $1
) limit $2;
