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


-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByUrl :one
SELECT * FROM feeds
WHERE url = $1;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows(
    id, created_at, updated_at, user_id, feed_id
  )
  Values(
    $1,
    $2,
    $3,
    $4,
    $5
  )
  RETURNING *
)

SELECT 
  inserted_feed_follow.*,
  feeds.name AS feed_name,
  users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id
INNER JOIN users ON users.id = inserted_feed_follow.user_id;



-- name: GetFeedFollowsForUser :many
SELECT
  ff.id,
  ff.created_at,
  ff.updated_at,
  ff.user_id,
  ff.feed_id,
  f.name AS feed_name,
  u.name AS user_name
FROM feed_follows AS ff
JOIN feeds AS f  ON f.id = ff.feed_id
JOIN users AS u  ON u.id = ff.user_id
WHERE ff.user_id = $1
ORDER BY ff.created_at DESC;   

-- name: DeleteFeedFollowByUserAndURL :exec
DELETE FROM feed_follows as ff
WHERE ff.user_id = $1
  AND ff.feed_id = (
    SELECT id FROM feeds WHERE url = $2
  );

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(),
    updated_at = NOW()
WHERE id = $1;


-- name: GetNextFeedToFetch :one
SELECT id, created_at, updated_at, name, url, user_id FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;


