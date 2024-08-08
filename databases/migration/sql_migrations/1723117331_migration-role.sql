-- +migrate Up
-- +migrate StatementBegin
create table
    role (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        created_at TIMESTAMP NOT NULL,
        created_by VARCHAR(255) NOT NULL,
        created_on VARCHAR(255) NOT NULL,
        modified_at TIMESTAMP NOT NULL,
        modified_by VARCHAR(255) NOT NULL,
        modified_on VARCHAR(255) NOT NULL
    );

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS role;

-- +migrate StatementEnd