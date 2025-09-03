-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO $$ BEGIN
    CREATE TYPE withdrawal_status AS ENUM ('pending', 'approved', 'rejected');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE withdrawal_requests (
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id           VARCHAR(255) NOT NULL,          -- external user ID
    requested_amt     DECIMAL(12,2) NOT NULL,
    processed_amt     DECIMAL(12,2) NULL,
    payment_method_id UUID NOT NULL,                  -- matches payment_methods.id
    status            withdrawal_status DEFAULT 'pending',
    reason            TEXT NULL,
    admin_id          UUID NULL,                       -- match admins.id if UUID
    txn_id            VARCHAR(100) NULL,
    txn_time          TIMESTAMP NULL,
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- no FK on user_id because it's a string
);

-- Indexes
CREATE INDEX idx_withdrawal_user_status ON withdrawal_requests(user_id, status);
CREATE INDEX idx_withdrawal_txn_id ON withdrawal_requests(txn_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS withdrawal_requests;
DROP TYPE IF EXISTS withdrawal_status;
-- +goose StatementEnd
