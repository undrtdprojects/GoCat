-- +migrate Up
-- +migrate StatementBegin
create table
    categories (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP NOT NULL,
        created_by VARCHAR(256) NOT NULL,
        modified_at TIMESTAMP NOT NULL,
        modified_by VARCHAR(256) NOT NULL
    );

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS categories;

-- +migrate StatementEnd