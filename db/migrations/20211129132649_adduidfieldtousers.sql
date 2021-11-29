-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN uid varchar(255) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN uid varchar(255) NOT NULL;
-- +goose StatementEnd
