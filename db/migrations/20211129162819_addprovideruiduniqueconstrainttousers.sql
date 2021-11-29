-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX provider_uid_idx ON users (provider, uid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX provider_uid_idx;
-- +goose StatementEnd
