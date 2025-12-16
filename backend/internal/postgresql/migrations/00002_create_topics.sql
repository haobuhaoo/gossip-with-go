-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Topics (
    topic_id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL NOT NULL,
    title TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Topics;
-- +goose StatementEnd
