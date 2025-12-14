-- name: ListTopics :many
SELECT * FROM Topics ORDER BY title;

-- name: FindTopicByID :one
SELECT * FROM Topics WHERE topic_id = $1;
