-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE user_cash_withdrawal (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id VARCHAR(50),
    requested_amt VARCHAR(50),
    reason VARCHAR(255),
    VPA_ID VARCHAR(55),
    status BOOLEAN DEFAULT FALSE,
    txn_id VARCHAR(55),
    txn_time VARCHAR(55),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

COMMENT ON COLUMN user_cash_withdrawal.status IS '0 means Pending, 1 means Transffered, 2 means Rejected';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_cash_withdrawal;
-- +goose StatementEnd

