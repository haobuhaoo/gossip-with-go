-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Posts (
    post_id BIGSERIAL PRIMARY KEY,
    topic_id BIGSERIAL NOT NULL,
    user_id BIGSERIAL NOT NULL,
    title TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES Users(user_id),
    FOREIGN KEY (topic_id) REFERENCES Topics(topic_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Posts;
-- +goose StatementEnd
