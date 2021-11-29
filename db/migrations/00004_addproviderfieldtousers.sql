-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN provider varchar(255) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN provider;
-- +goose StatementEnd
