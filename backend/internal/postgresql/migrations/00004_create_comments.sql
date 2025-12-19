-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Comments (
    comment_id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL NOT NULL,
    post_id BIGSERIAL NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES Users(user_id),
    FOREIGN KEY (post_id) REFERENCES Posts(post_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Comments;
-- +goose StatementEnd
