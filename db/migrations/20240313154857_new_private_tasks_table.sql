-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS private_tasks
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    description TEXT,
    status BOOLEAN NOT NULL DEFAULT false,
    date_create TIMESTAMP NOT NULL DEFAULT now(),
    deadline TIMESTAMP,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    board_id INT REFERENCES private_boards(id) ON DELETE CASCADE,
    category_id INT REFERENCES private_categories(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS private_tasks;
-- +goose StatementEnd
