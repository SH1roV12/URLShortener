-- +goose Up
-- +goose StatementBegin
CREATE TABLE links(
    ID int AUTO_INCREMENT PRIMARY KEY,
    OriginalURL text not null,
    ShortURL VARCHAR(255) not null UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS links;
-- +goose StatementEnd
