-- +goose Up
ALTER TABLE hta.media ADD COLUMN sys_status varchar(32) DEFAULT 'active';

-- +goose Down
ALTER TABLE hta.media DROP COLUMN sys_status;
