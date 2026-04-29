-- +goose Up
CREATE TABLE IF NOT EXISTS hta.user (
    id text PRIMARY KEY,
    created_at timestamp without time zone NOT NULL DEFAULT LOCALTIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT LOCALTIMESTAMP,
    deleted_at timestamp without time zone,
    first_name varchar(255),
    last_name varchar(255),
    email varchar(255) UNIQUE,
    picture text
);

CREATE INDEX IF NOT EXISTS idx_user_deleted_at ON hta.user (deleted_at);

-- +goose Down
DROP TABLE IF EXISTS hta.user;
