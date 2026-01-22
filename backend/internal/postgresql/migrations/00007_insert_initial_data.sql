-- +goose Up
-- +goose StatementBegin
INSERT INTO Users (name) VALUES ('admin'), ('cvwo'), ('tester');

INSERT INTO Topics (user_id, title) VALUES (1, 'food'), (1, 'sports'), (2, 'tech'), (3, 'health');

INSERT INTO Posts (topic_id, user_id, title, description) VALUES
(1, 1, 'Best street food in Singapore', 'What are your favorite street food stalls or hawker centers in Singapore?'),
(1, 2, 'Homemade pasta tips', 'Any tips to improve the texture and flavor of homemade pasta?'),
(1, 3, 'Healthy dessert ideas', 'Looking for dessert ideas that are tasty but not too high in sugar.');

INSERT INTO Comments (user_id, post_id, description) VALUES
-- Comments for Post 1
(2, 1, 'Maxwell Food Centre is a must-visit, especially for chicken rice.'),
(3, 1, 'Old Airport Road has a great variety and very reasonable prices.'),
-- Comments for Post 2
(1, 2, 'Use 00 flour if you can and let the dough rest longer before rolling.'),
(3, 2, 'Fresh eggs make a big difference in both color and taste.'),
-- Comments for Post 3
(1, 3, 'Greek yogurt with honey and berries works great for me.'),
(2, 3, 'Dark chocolate with a high cocoa percentage can satisfy sweet cravings.');

INSERT INTO Post_Votes (post_id, user_id, vote) VALUES
(1, 1, 1), (1, 2, 1), (1, 3, -1), (2, 2, -1), (2, 3, -1), (3, 1, 1), (3, 2, 1), (3, 3, 1);

INSERT INTO Comment_Votes (comment_id, user_id, vote) VALUES
(1, 1, 1), (1, 2, 1), (2, 2, -1), (2, 3, -1), (3, 1, 1), (3, 2, 1), (3, 3, 1), (5, 3, -1), (6, 1, 1), (6, 2, -1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
