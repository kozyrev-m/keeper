-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    login VARCHAR UNIQUE,
    encrypted_password VARCHAR
);
CREATE TABLE private_data (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    owner_id INT NOT NULL,
    type_id INT NOT NULL,
    metadata VARCHAR NOT NULL,
    content VARCHAR NOT NULL,
    created_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE private_data;
-- +goose StatementEnd
