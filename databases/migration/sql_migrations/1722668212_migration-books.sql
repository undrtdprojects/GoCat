-- +migrate Up
-- +migrate StatementBegin
create table
    books (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description VARCHAR(255) NOT NULL,
        image_url VARCHAR(255) NOT NULL,
        release_year INT NOT NULL,
        price INT NOT NULL,
        total_page INT NOT NULL,
        thickness VARCHAR(255) NOT NULL,
        category_id INT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        created_by VARCHAR(256) NOT NULL,
        modified_at TIMESTAMP NOT NULL,
        modified_by VARCHAR(256) NOT NULL
    );

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS books;

-- +migrate StatementEnd