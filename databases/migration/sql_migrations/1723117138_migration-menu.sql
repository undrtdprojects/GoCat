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

    INSERT INTO menu (id, name, price, category_id, created_at, created_by, created_on, modified_at, modified_by, modified_on)
    VALUES 
    ('MAC-0001', 'Spaghetti Carbonara', 12000, 'MAC', CURRENT_TIMESTAMP, 'admin', 'system', CURRENT_TIMESTAMP, 'admin', 'system'),
    ('MAC-0002', 'Margherita Pizza', 15000, 'MAC', CURRENT_TIMESTAMP, 'admin', 'system', CURRENT_TIMESTAMP, 'admin', 'system'),
    ('APT-0001', 'Caesar Salad', 8000, 'APT', CURRENT_TIMESTAMP, 'admin', 'system', CURRENT_TIMESTAMP, 'admin', 'system'),
    ('BEV-0001', 'Iced Latte', 5000, 'BEV', CURRENT_TIMESTAMP, 'admin', 'system', CURRENT_TIMESTAMP, 'admin', 'system'),
    ('DST-0001', 'Chocolate Brownie', 7000, 'DST', CURRENT_TIMESTAMP, 'admin', 'system', CURRENT_TIMESTAMP, 'admin', 'system');

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS menu;

-- +migrate StatementEnd