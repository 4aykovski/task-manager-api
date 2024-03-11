-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS refresh_sessions
(
    token       text PRIMARY KEY,
    user_id     uuid      NOT NULL,
    fingerprint text      NOT NULL,
    created_at  timestamp NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS refresh_sessions;
-- +goose StatementEnd
