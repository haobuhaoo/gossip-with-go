-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Comment_Votes (
    comment_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    vote SMALLINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT Comment_Votes_pk PRIMARY KEY (comment_id, user_id),
    CONSTRAINT vote_valid CHECK (vote in (-1, 1)),
    FOREIGN KEY (comment_id) REFERENCES Comments(comment_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Comment_Votes;
-- +goose StatementEnd
