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

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $2, updated_at = $2
WHERE id = $1;

-- name: GetNextFeedToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST;
