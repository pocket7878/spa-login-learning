-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	id BIGINT NOT NULL,
	email varchar(255) NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users
-- +goose StatementEnd
