-- +goose Up
-- +goose StatementBegin
ALTER TABLE devices ADD COLUMN email text NOT NULL DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE devices DROP COLUMN email;
-- +goose StatementEnd
