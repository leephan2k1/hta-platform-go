-- +goose Up
-- Indexes for media_chapter to speed up retrieval and ordering
CREATE INDEX IF NOT EXISTS idx_media_chapter_media_id_order ON hta.media_chapter (media_id, "order" DESC) WHERE deleted_at IS NULL;

-- Index for images to speed up retrieval by resource_id (media_id)
CREATE INDEX IF NOT EXISTS idx_image_resource_id ON hta.image (resource_id) WHERE deleted_at IS NULL;

-- Index for media_other_name
CREATE INDEX IF NOT EXISTS idx_media_other_name_media_id ON hta.media_other_name (media_id) WHERE deleted_at IS NULL;

-- Indexes for join tables
CREATE INDEX IF NOT EXISTS idx_media_to_category_media_id ON hta.media_to_category (media_id);
CREATE INDEX IF NOT EXISTS idx_media_to_author_media_id ON hta.media_to_author (media_id);

-- +goose Down
DROP INDEX IF EXISTS hta.idx_media_chapter_media_id_order;
DROP INDEX IF EXISTS hta.idx_image_resource_id;
DROP INDEX IF EXISTS hta.idx_media_other_name_media_id;
DROP INDEX IF EXISTS hta.idx_media_to_category_media_id;
DROP INDEX IF EXISTS hta.idx_media_to_author_media_id;
