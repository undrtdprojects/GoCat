-- +migrate Up
-- +migrate StatementBegin
create table
    categories (
        id VARCHAR(3) PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP NOT NULL,
        created_by VARCHAR(255) NOT NULL,
        created_on VARCHAR(255) NOT NULL,
        modified_at TIMESTAMP NOT NULL,
        modified_by VARCHAR(255) NOT NULL,
        modified_on VARCHAR(255) NOT NULL
    );


    INSERT INTO categories 
            (id, name, created_at, created_by, created_on, modified_at, modified_by, modified_on)
        VALUES 
            ('APT', 'Appetizers', CURRENT_TIMESTAMP, 'System', 'Initial Setup', CURRENT_TIMESTAMP, 'System', 'Initial Setup'),
            ('MAC', 'Main Course', CURRENT_TIMESTAMP, 'System', 'Initial Setup', CURRENT_TIMESTAMP, 'System', 'Initial Setup'),
            ('DST', 'Desserts', CURRENT_TIMESTAMP, 'System', 'Initial Setup', CURRENT_TIMESTAMP, 'System', 'Initial Setup'),
            ('BEV', 'Beverages', CURRENT_TIMESTAMP, 'System', 'Initial Setup', CURRENT_TIMESTAMP, 'System', 'Initial Setup'),
            ('SPC', 'Specials', CURRENT_TIMESTAMP, 'System', 'Initial Setup', CURRENT_TIMESTAMP, 'System', 'Initial Setup');

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS categories;

-- +migrate StatementEnd