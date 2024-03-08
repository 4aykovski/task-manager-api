-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id            UUID PRIMARY KEY,
    login         VARCHAR(128) UNIQUE NOT NULL,
    email         VARCHAR(128) UNIQUE NOT NULL,
    password      VARCHAR(60)         NOT NULL,
    registered_at TIMESTAMP           NOT NULL DEFAULT NOW(),
    last_seen_at  TIMESTAMP           NOT NULL DEFAULT NOW(),
    is_admin      BOOLEAN             NOT NULL DEFAULT FALSE,
    language      VARCHAR(2)          NOT NULL,
    theme         VARCHAR(10)         NOT NULL DEFAULT 'light',
    about         TEXT                NOT NULL DEFAULT ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
