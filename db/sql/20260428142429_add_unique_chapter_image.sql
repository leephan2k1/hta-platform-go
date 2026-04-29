-- +goose Up
-- +goose StatementBegin
ALTER TABLE hta.chapter_image ADD CONSTRAINT unique_order_chapter UNIQUE ("order", chapter_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE hta.chapter_image DROP CONSTRAINT IF EXISTS unique_order_chapter;
-- +goose StatementEnd
