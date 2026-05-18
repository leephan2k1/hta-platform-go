-- +goose Up
-- +goose StatementBegin
CREATE TABLE "hta"."user_reading_progress" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    "deleted_at" timestamptz,
    "user_id" text NOT NULL,
    "chapter_id" uuid NOT NULL,
    "media_id" uuid NOT NULL,
    "chapter_image_order" integer
);

-- B-tree index for fast look up on (user_id, chapter_id, media_id)
CREATE INDEX "idx_user_reading_progress_lookup" ON "hta"."user_reading_progress" ("user_id", "chapter_id", "media_id");

-- Unique index for (user_id, chapter_id, media_id, chapter_image_order)
CREATE UNIQUE INDEX "idx_user_reading_progress_unique" ON "hta"."user_reading_progress" ("user_id", "chapter_id", "media_id", "chapter_image_order");

CREATE TABLE "hta"."user_reading_sessions" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz NOT NULL DEFAULT now(),
    "deleted_at" timestamptz,
    "user_id" text NOT NULL,
    "chapter_id" uuid NOT NULL,
    "media_id" uuid NOT NULL,
    "started_at" timestamptz NOT NULL,
    "ended_at" timestamptz,
    "duration" bigint
);

-- B-tree index on columns (user_id, chapter_id, media_id)
CREATE INDEX "idx_user_reading_sessions_lookup" ON "hta"."user_reading_sessions" ("user_id", "chapter_id", "media_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "hta"."user_reading_sessions";
DROP TABLE IF EXISTS "hta"."user_reading_progress";
-- +goose StatementEnd
