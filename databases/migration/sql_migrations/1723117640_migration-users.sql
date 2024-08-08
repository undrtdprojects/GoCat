-- +migrate Up
-- +migrate StatementBegin
create table
    users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) UNIQUE,
        password VARCHAR(255) NOT NULL,
        role_id INT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        created_by VARCHAR(255),
        created_on VARCHAR(255) NOT NULL,
        modified_at TIMESTAMP NOT NULL,
        modified_by VARCHAR(255),
        modified_on VARCHAR(255) NOT NULL
    );

-- +migrate StatementEnd
-- +migrate Down
-- +migrate StatementBegin
DROP TABLE IF EXISTS users;

-- +migrate StatementEnd