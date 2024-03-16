-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS project_tasks
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    description TEXT,
    status BOOLEAN NOT NULL DEFAULT false,
    date_create TIMESTAMP NOT NULL DEFAULT now(),
    deadline TIMESTAMP,
    project_id INT REFERENCES projects(id) ON DELETE CASCADE,
    board_id INT REFERENCES project_boards(id) ON DELETE CASCADE,
    category_id INT REFERENCES project_categories(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS project_tasks;
-- +goose StatementEnd
