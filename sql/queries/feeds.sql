-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE feeds.url = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: ResetFeeds :exec
DELETE FROM feeds;
