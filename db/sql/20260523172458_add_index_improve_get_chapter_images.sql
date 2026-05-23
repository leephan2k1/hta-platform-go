-- +goose Up
CREATE INDEX IF NOT EXISTS idx_chapter_image_chapter_id
    ON hta.chapter_image USING btree (chapter_id ASC)
    WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_media_chapter_media_id_url
    ON hta.media_chapter USING btree (media_id ASC, url ASC)
    WHERE deleted_at IS NULL;

-- +goose Down
DROP INDEX IF EXISTS hta.idx_chapter_image_chapter_id;
DROP INDEX IF EXISTS hta.idx_media_chapter_media_id_url;
