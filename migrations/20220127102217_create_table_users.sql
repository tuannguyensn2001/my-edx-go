-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users
-- +goose StatementEnd
