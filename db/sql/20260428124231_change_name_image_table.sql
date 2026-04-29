-- +goose Up
-- +goose StatementBegin
ALTER TABLE hta.images RENAME TO image;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE hta.image RENAME TO images;
-- +goose StatementEnd
