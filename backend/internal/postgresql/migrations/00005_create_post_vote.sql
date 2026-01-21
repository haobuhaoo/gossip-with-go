-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Post_Votes (
    post_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    vote SMALLINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT Post_Votes_pk PRIMARY KEY (post_id, user_id),
    CONSTRAINT vote_valid CHECK (vote in (-1, 1)),
    FOREIGN KEY (post_id) REFERENCES Posts(post_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Post_Votes;
-- +goose StatementEnd
