-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD UNIQUE(id)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
