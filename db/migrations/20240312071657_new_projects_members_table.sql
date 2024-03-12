-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS projects_members
(
    project_id INT  NOT NULL REFERENCES projects (id) ON DELETE CASCADE,
    user_id    UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    status     INT  NOT NULL,
    UNIQUE (project_id, user_id),
    PRIMARY KEY (project_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS projects_members;
-- +goose StatementEnd
