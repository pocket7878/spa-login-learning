-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	provider varchar(255) NOT NULL,
	uid varchar(255) NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users
-- +goose StatementEnd
