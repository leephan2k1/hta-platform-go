-- +goose Up
-- +goose StatementBegin
ALTER TABLE hta.chapter_image DROP COLUMN IF EXISTS url;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE hta.chapter_image ADD COLUMN url text;
-- +goose StatementEnd
