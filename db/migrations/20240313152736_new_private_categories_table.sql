-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS private_categories
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(128) NOT NULL,
    color      VARCHAR(128) NOT NULL,
    project_id INTEGER      NOT NULL REFERENCES projects (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS private_categories;
-- +goose StatementEnd
