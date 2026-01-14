-- Users Queries
-- name: FindUserByName :one
SELECT * FROM Users WHERE name = $1;

-- name: CreateUser :one
INSERT INTO Users (name) VALUES ($1) RETURNING *;

-- Topics Queries
-- name: ListTopics :many
SELECT * FROM Topics ORDER BY title;

-- name: FindTopicByID :one
SELECT * FROM Topics WHERE topic_id = $1;

-- name: CreateTopic :one
INSERT INTO Topics (user_id, title) VALUES ($1, $2) RETURNING *;

-- name: UpdateTopic :one
UPDATE Topics SET title = $2 WHERE topic_id = $1 RETURNING *;

-- name: DeleteTopic :execrows
DELETE FROM Topics WHERE topic_id = $1;

-- Posts Queries
-- name: FindPostsByTopic :many
SELECT p.post_id, p.topic_id, p.user_id, u.name AS username, p.title, p.description, p.created_at, p.updated_at
FROM Posts p JOIN Users u ON u.user_id = p.user_id WHERE p.topic_id = $1 ORDER BY p.updated_at DESC;

-- name: FindPostByID :one
SELECT p.post_id, p.topic_id, p.user_id, u.name AS username, p.title, p.description, p.created_at, p.updated_at
FROM Posts p JOIN Users u ON u.user_id = p.user_id WHERE post_id = $1;

-- name: CreatePost :one
INSERT INTO Posts (topic_id, user_id, title, description) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdatePost :one
UPDATE Posts SET title = $2, description = $3, updated_at = now() WHERE post_id = $1 RETURNING *;

-- name: UpdatePostStatus :one
UPDATE Posts SET updated_at = now() WHERE post_id = $1 RETURNING *;

-- name: DeletePost :execrows
DELETE FROM Posts WHERE post_id = $1;

-- Comments Queries
-- name: FindCommentsByPost :many
SELECT c.comment_id, c.user_id, u.name as username, c.post_id, c.description, c.created_at, c.updated_at
FROM Comments c JOIN Users u ON u.user_id = c.user_id WHERE c.post_id = $1 ORDER BY c.updated_at DESC;

-- name: CreateComment :one
INSERT INTO Comments (user_id, post_id, description) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateComment :one
UPDATE Comments SET description = $3, updated_at = now() WHERE comment_id = $1 AND post_id = $2 RETURNING *;

-- name: DeleteComment :execrows
DELETE FROM Comments WHERE comment_id = $1;
