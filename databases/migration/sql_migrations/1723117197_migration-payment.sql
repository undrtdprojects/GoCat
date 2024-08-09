-- +migrate Up
-- +migrate StatementBegin
create table
    payment (
        id SERIAL PRIMARY KEY,
        name VARCHAR(50) NOT NULL,
        created_at TIMESTAMP NOT NULL,
        created_by VARCHAR(255) NOT NULL,
        created_on VARCHAR(255) NOT NULL,
        modified_at TIMESTAMP NOT NULL,
        modified_by VARCHAR(255) NOT NULL,
        modified_on VARCHAR(255) NOT NULL
    );

    INSERT INTO payment (name, created_at, created_by, created_on, modified_at, modified_by, modified_on)
    VALUES 
    ('Cash', CURRENT_TIMESTAMP, 'System', 'Initial Setup', CURRENT_TIMESTAMP, 'System', 'Initial Setup'),
    ('Credit Card', CURRENT_TIMESTAMP, 'System', 'Initial Setup', CURRENT_TIMESTAMP, 'System', 'Initial Setup'),
    ('Debit Card', CURRENT_TIMESTAMP, 'System', 'Initial Setup', CURRENT_TIMESTAMP, 'System', 'Initial Setup'),
    ('PayPal', CURRENT_TIMESTAMP, 'System', 'Initial Setup', CURRENT_TIMESTAMP, 'System', 'Initial Setup'),
    ('Bank Transfer', CURRENT_TIMESTAMP, 'System', 'Initial Setup', CURRENT_TIMESTAMP, 'System', 'Initial Setup');

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS payment;

-- +migrate StatementEnd