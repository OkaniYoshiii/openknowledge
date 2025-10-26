-- +goose Up
CREATE TABLE posts (
    id INT UNSIGNED NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL UNIQUE,
    content TEXT NOT NULL
);

-- +goose Down
DROP TABLE posts;
