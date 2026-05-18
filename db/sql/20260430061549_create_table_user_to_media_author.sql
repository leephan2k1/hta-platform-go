-- +goose Up
CREATE TABLE IF NOT EXISTS hta.user_to_author (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at timestamp without time zone NOT NULL DEFAULT LOCALTIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT LOCALTIMESTAMP,
    deleted_at timestamp without time zone,
    user_id text NOT NULL,
    author_id uuid NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES hta.user(id) ON DELETE CASCADE,
    CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES hta.author(id) ON DELETE CASCADE,
    UNIQUE(user_id, author_id)
);

CREATE INDEX IF NOT EXISTS idx_user_to_author_deleted_at ON hta.user_to_author (deleted_at);

CREATE TABLE IF NOT EXISTS hta.user_to_media (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at timestamp without time zone NOT NULL DEFAULT LOCALTIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT LOCALTIMESTAMP,
    deleted_at timestamp without time zone,
    user_id text NOT NULL,
    media_id uuid NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES hta.user(id) ON DELETE CASCADE,
    CONSTRAINT fk_media FOREIGN KEY (media_id) REFERENCES hta.media(id) ON DELETE CASCADE,
    UNIQUE(user_id, media_id)
);

CREATE INDEX IF NOT EXISTS idx_user_to_media_deleted_at ON hta.user_to_media (deleted_at);

-- +goose Down
DROP TABLE IF EXISTS hta.user_to_author;
DROP TABLE IF EXISTS hta.user_to_media;
