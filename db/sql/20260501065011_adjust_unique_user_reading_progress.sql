-- +goose Up
-- +goose StatementBegin
-- Drop the old unique index
DROP INDEX IF EXISTS "hta"."idx_user_reading_progress_unique";

-- Create new unique index on (user_id, media_id)
CREATE UNIQUE INDEX "idx_user_reading_progress_unique" ON "hta"."user_reading_progress" ("user_id", "media_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop the new unique index
DROP INDEX IF EXISTS "hta"."idx_user_reading_progress_unique";

-- Restore the old unique index
CREATE UNIQUE INDEX "idx_user_reading_progress_unique" ON "hta"."user_reading_progress" ("user_id", "chapter_id", "media_id", "chapter_image_order");
-- +goose StatementEnd
