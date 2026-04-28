-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS hta.images (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    created_at timestamp without time zone NOT NULL DEFAULT LOCALTIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT LOCALTIMESTAMP,
    deleted_at timestamp without time zone,
    url text NOT NULL UNIQUE,
    description text,
    resource_id uuid NOT NULL,
    source varchar(125),
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS hta.images;
-- +goose StatementEnd
