-- +goose Up
CREATE TABLE users_goose (
    id int NOT NULL PRIMARY KEY,
    username text,
    name text,
    surname text
);

INSERT INTO users_goose VALUES
(0, 'root', '', ''),
(1, 'goose', "goose insert", "goose surname");

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE users_goose;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
