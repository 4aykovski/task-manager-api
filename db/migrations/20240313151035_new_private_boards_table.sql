-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS private_boards
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    color VARCHAR(128) NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS private_boards;
-- +goose StatementEnd
