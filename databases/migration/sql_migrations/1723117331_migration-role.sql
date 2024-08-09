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

    INSERT INTO role (name, created_at, created_by, created_on, modified_at, modified_by, modified_on)
    VALUES 
    ('Admin', CURRENT_TIMESTAMP, 'System', 'Initial Setup', CURRENT_TIMESTAMP, 'System', 'Initial Setup'),
    ('User', CURRENT_TIMESTAMP, 'System', 'Initial Setup', CURRENT_TIMESTAMP, 'System', 'Initial Setup'),
    ('Manager', CURRENT_TIMESTAMP, 'System', 'Initial Setup', CURRENT_TIMESTAMP, 'System', 'Initial Setup');



    This

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS role;

-- +migrate StatementEnd