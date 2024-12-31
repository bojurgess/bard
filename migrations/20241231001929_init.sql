-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id TEXT PRIMARY KEY,
    display_name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS tokens(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT NOT NULL UNIQUE REFERENCES users(id),
    access_token TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    expires_at DATETIME NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users";
DROP TABLE "tokens";
-- +goose StatementEnd
