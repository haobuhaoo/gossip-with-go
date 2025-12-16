-- name: ListTopics :many
SELECT * FROM Topics ORDER BY title;

-- name: FindTopicByID :one
SELECT * FROM Topics WHERE topic_id = $1;

-- name: CreateTopic :one
INSERT INTO Topics (user_iD, title) VALUES ($1, $2) RETURNING *;

-- name: UpdateTopic :one
UPDATE Topics SET title = $2 WHERE topic_id = $1 RETURNING *;

-- name: DeleteTopic :execrows
DELETE FROM Topics WHERE topic_id = $1;