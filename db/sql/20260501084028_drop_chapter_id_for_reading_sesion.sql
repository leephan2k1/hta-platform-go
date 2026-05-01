-- +goose Up
-- +goose StatementBegin
DROP INDEX IF EXISTS "hta"."idx_user_reading_sessions_lookup";
ALTER TABLE "hta"."user_reading_sessions" DROP COLUMN IF EXISTS "chapter_id";
CREATE INDEX "idx_user_reading_sessions_lookup" ON "hta"."user_reading_sessions" ("user_id", "media_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "hta"."idx_user_reading_sessions_lookup";
ALTER TABLE "hta"."user_reading_sessions" ADD COLUMN "chapter_id" uuid;
CREATE INDEX "idx_user_reading_sessions_lookup" ON "hta"."user_reading_sessions" ("user_id", "chapter_id", "media_id");
-- +goose StatementEnd
