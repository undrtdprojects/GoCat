-- +migrate Up
-- +migrate StatementBegin
create table
    transaction1 (
        id SERIAL PRIMARY KEY,
        transaction_id VARCHAR(20) NOT NULL,
        menu_id VARCHAR(20) NOT NULL,
        date_transaction TIMESTAMP NOT NULL,
        qty INT NOT NULL,
        total_price INT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        created_by VARCHAR(255) NOT NULL,
        created_on VARCHAR(255) NOT NULL,
        modified_at TIMESTAMP NOT NULL,
        modified_by VARCHAR(255) NOT NULL,
        modified_on VARCHAR(255) NOT NULL
    );

    INSERT INTO transaction1 (
        transaction_id, 
        menu_id, 
        date_transaction, 
        qty, 
        total_price, 
        created_at, 
        created_by, 
        created_on, 
        modified_at, 
        modified_by, 
        modified_on
    ) VALUES (
        'CAT-00001',
        'BEV-0001',
        CURRENT_TIMESTAMP,
        10,
        50000,
        CURRENT_TIMESTAMP,
        'system',
        'localhost',
        CURRENT_TIMESTAMP,
        'system',
        'localhost'
    );

-- +migrate StatementEnd
-- +migrate Down
-- +migrate StatementBegin
DROP TABLE IF EXISTS transaction1;

-- +migrate StatementEnd