-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE miniapp_transactions (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    campaign_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    commission VARCHAR(50),
    user_commission VARCHAR(50),
    order_id VARCHAR(50),
    reference_id VARCHAR(50) NOT NULL,
    sale_amount DECIMAL(10, 2),
    status VARCHAR(50),
    subid VARCHAR(255),
    subid1 VARCHAR(255),
    subid2 VARCHAR(255),
    miniapp_id VARCHAR(255),
    commission_percentage VARCHAR(25),
    transaction_date TIMESTAMP,
    transaction_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN miniapp_transactions.status IS '0 means pending, 1 means verified, 2 means rejected';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS miniapp_transactions;
-- +goose StatementEnd
