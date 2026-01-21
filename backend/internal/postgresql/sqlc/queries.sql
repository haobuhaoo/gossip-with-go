-- Users Queries
-- name: FindUserByName :one
SELECT * FROM Users WHERE name = $1;

-- name: FindUserByID :one
SELECT * FROM Users WHERE user_id = $1;

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
UPDATE Topics SET title = $3 WHERE topic_id = $1 AND user_id = $2 RETURNING *;

-- name: DeleteTopic :execrows
DELETE FROM Topics WHERE topic_id = $1 AND user_id = $2;

-- name: SearchTopic :many
SELECT * FROM Topics WHERE title ILIKE '%' || $1 ||'%' ORDER BY title;

-- Posts Queries
-- name: FindPostsByTopic :many
SELECT DISTINCT p.post_id, p.topic_id, p.user_id, u.name AS username,
p.title, p.description, p.created_at, p.updated_at,
COUNT(v.vote) FILTER (WHERE v.vote = 1) OVER (PARTITION BY p.post_id) AS likes,
COUNT(v.vote) FILTER (WHERE v.vote = -1) OVER (PARTITION BY p.post_id) AS dislikes,
MAX(uv.vote) OVER (PARTITION BY p.post_id) AS user_vote
FROM Posts p
JOIN Users u ON u.user_id = p.user_id
LEFT JOIN Post_Votes v ON p.post_id = v.post_id
LEFT JOIN Post_Votes uv ON p.post_id = uv.post_id AND uv.user_id = $2
WHERE p.topic_id = $1
ORDER BY likes DESC, p.updated_at DESC;

-- name: FindPostByID :one
SELECT DISTINCT p.post_id, p.topic_id, p.user_id, u.name AS username,
p.title, p.description, p.created_at, p.updated_at,
COUNT(v.vote) FILTER (WHERE v.vote = 1) OVER (PARTITION BY p.post_id) AS likes,
COUNT(v.vote) FILTER (WHERE v.vote = -1) OVER (PARTITION BY p.post_id) AS dislikes,
MAX(uv.vote) OVER (PARTITION BY p.post_id) AS user_vote
FROM Posts p
JOIN Users u ON u.user_id = p.user_id
LEFT JOIN Post_Votes v ON p.post_id = v.post_id
LEFT JOIN Post_Votes uv ON p.post_id = uv.post_id AND uv.user_id = $3
WHERE p.post_id = $1 AND p.topic_id = $2;

-- name: CreatePost :one
INSERT INTO Posts (topic_id, user_id, title, description) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdatePost :one
UPDATE Posts SET title = $3, description = $4, updated_at = now() WHERE post_id = $1 AND user_id = $2 RETURNING *;

-- name: UpdatePostStatus :one
UPDATE Posts SET updated_at = now() WHERE post_id = $1 AND user_id = $2 RETURNING *;

-- name: DeletePost :execrows
DELETE FROM Posts WHERE post_id = $1 AND user_id = $2;

-- name: SearchPost :many
SELECT DISTINCT p.post_id, p.topic_id, p.user_id, u.name AS username,
p.title, p.description, p.created_at, p.updated_at,
COUNT(v.vote) FILTER (WHERE v.vote = 1) OVER (PARTITION BY p.post_id)AS likes,
COUNT(v.vote) FILTER (WHERE v.vote = -1) OVER (PARTITION BY p.post_id)AS dislikes,
MAX(uv.vote) OVER (PARTITION BY p.post_id)AS user_vote
FROM Posts p
JOIN Users u ON u.user_id = p.user_id
LEFT JOIN Post_Votes v ON p.post_id = v.post_id
LEFT JOIN Post_Votes uv ON p.post_id = uv.post_id AND uv.user_id = $3
WHERE p.topic_id = $1 AND (p.title ILIKE '%' || $2 ||'%' OR p.description ILIKE '%' || $2 ||'%')
ORDER BY likes DESC, p.updated_at DESC;

-- Comments Queries
-- name: FindCommentsByPost :many
SELECT c.comment_id, c.user_id, u.name as username, c.post_id, c.description, c.created_at, c.updated_at
FROM Comments c JOIN Users u ON u.user_id = c.user_id WHERE c.post_id = $1 ORDER BY c.updated_at DESC;

-- name: CreateComment :one
INSERT INTO Comments (user_id, post_id, description) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateComment :one
UPDATE Comments SET description = $4, updated_at = now()
WHERE comment_id = $1 AND post_id = $2 AND user_id = $3 RETURNING *;

-- name: DeleteComment :execrows
DELETE FROM Comments WHERE comment_id = $1 AND user_id = $2;

-- Post Votes
-- name: LikesPost :one
INSERT INTO Post_Votes (post_id, user_id, vote) VALUES ($1, $2, 1)
ON CONFLICT (post_id, user_id) DO UPDATE SET vote = 1 WHERE Post_Votes.vote <> 1 RETURNING *;

-- name: DislikesPost :one
INSERT INTO Post_Votes (post_id, user_id, vote) VALUES ($1, $2, -1)
ON CONFLICT (post_id, user_id) DO UPDATE SET vote = -1 WHERE Post_Votes.vote <> -1 RETURNING *;

-- name: RemovePostVote :execrows
DELETE FROM Post_Votes WHERE post_id = $1 AND user_id = $2;

-- Comment Votes
-- name: LikesComment :one
INSERT INTO Comment_Votes (comment_id, user_id, vote) VALUES ($1, $2, 1)
ON CONFLICT (comment_id, user_id) DO UPDATE SET vote = 1 WHERE Comment_Votes.vote <> 1 RETURNING *;

-- name: DislikesComment :one
INSERT INTO Comment_Votes (comment_id, user_id, vote) VALUES ($1, $2, -1)
ON CONFLICT (comment_id, user_id) DO UPDATE SET vote = -1 WHERE Comment_Votes.vote <> -1 RETURNING *;

-- name: RemoveCommentVote :execrows
DELETE FROM Comment_Votes WHERE comment_id = $1 AND user_id = $2;

-- name: CountVote :one
SELECT COUNT(*) FILTER (WHERE vote = 1) AS likes, COUNT(*) FILTER (WHERE vote = -1) AS dislikes
FROM Comment_Votes WHERE comment_id = $1;
