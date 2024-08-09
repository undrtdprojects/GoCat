-- +migrate Up
-- +migrate StatementBegin
create table
    transaction0 (
        id VARCHAR(20) PRIMARY KEY,
        user_id INT NOT NULL,
        payment_id INT NOT NULL,
        grand_total_price INT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        created_by VARCHAR(255) NOT NULL,
        created_on VARCHAR(255) NOT NULL,
        modified_at TIMESTAMP NOT NULL,
        modified_by VARCHAR(255) NOT NULL,
        modified_on VARCHAR(255) NOT NULL
    );

    INSERT INTO transaction0 (
        id, 
        user_id, 
        payment_id, 
        grand_total_price, 
        created_at, 
        created_by, 
        created_on, 
        modified_at, 
        modified_by, 
        modified_on
    ) VALUES (
        'CAT-00001',
        1,
        1,
        50000,
        CURRENT_TIMESTAMP,
        'system',
        'web',
        CURRENT_TIMESTAMP,
        'system',
        'web'
    );

-- +migrate StatementEnd
-- +migrate Down
-- +migrate StatementBegin
DROP TABLE IF EXISTS transaction0;

-- +migrate StatementEnd