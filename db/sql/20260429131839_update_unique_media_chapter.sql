-- +goose Up
ALTER TABLE hta.media_chapter DROP CONSTRAINT IF EXISTS media_chapter_name_unique;
ALTER TABLE hta.media_chapter DROP CONSTRAINT IF EXISTS media_chapter_url_unique;
CREATE UNIQUE INDEX idx_media_chapter_unique ON hta.media_chapter (media_id, url, "order");

-- +goose Down
DROP INDEX IF EXISTS hta.idx_media_chapter_unique;
ALTER TABLE hta.media_chapter ADD CONSTRAINT media_chapter_name_unique UNIQUE (name);
ALTER TABLE hta.media_chapter ADD CONSTRAINT media_chapter_url_unique UNIQUE (url);

