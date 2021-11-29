-- +goose Up
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN email;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN email varchar(255) NOT NULL;
-- +goose StatementEnd
