-- +goose Up
-- +goose StatementBegin
CREATE TABLE todos(
	id serial PRIMARY KEY,
	user_id BIGINT NOT NULL,
	description text,
	FOREIGN KEY (user_id) REFERENCES users(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todos
-- +goose StatementEnd
