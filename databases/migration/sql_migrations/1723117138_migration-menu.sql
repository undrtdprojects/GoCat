-- +migrate Up
-- +migrate StatementBegin
create table
    menu (
        id VARCHAR(8) PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        price INT NOT NULL,
        category_id VARCHAR(20) NOT NULL,
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

DROP TABLE IF EXISTS menu;

-- +migrate StatementEnd