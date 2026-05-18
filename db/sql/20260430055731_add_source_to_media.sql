-- +goose Up
ALTER TABLE hta.media ADD COLUMN IF NOT EXISTS source varchar(125);

-- +goose Down
ALTER TABLE hta.media DROP COLUMN IF EXISTS source;
