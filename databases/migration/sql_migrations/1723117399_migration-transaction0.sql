-- +migrate Up
-- +migrate StatementBegin
create table
    transaction0 (
        id VARCHAR(20) PRIMARY KEY,
        user_id INT NOT NULL,
        grand_total_price INT NOT NULL,
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
DROP TABLE IF EXISTS transaction0;

-- +migrate StatementEnd