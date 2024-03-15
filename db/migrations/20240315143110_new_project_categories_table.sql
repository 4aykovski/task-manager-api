-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS project_categories
(
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(128) NOT NULL,
    color    VARCHAR(128) NOT NULL,
    board_id INTEGER      NOT NULL REFERENCES project_boards (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS project_categories;
-- +goose StatementEnd
