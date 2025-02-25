-- +goose Up
-- +goose StatementBegin
COPY tasks (title, description)
    FROM '/migrations/todos.txt'
    WITH (FORMAT csv, DELIMITER ',', QUOTE '"',HEADER false);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
