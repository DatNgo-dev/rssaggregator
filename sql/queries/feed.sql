-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds where id = $1 AND user_id = $2;

-- name: GetFeeds :many
SELECT * FROM feeds where user_id = $1;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetch :one
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
where id = $1
RETURNING *;
