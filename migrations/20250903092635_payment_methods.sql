-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO $$ BEGIN
    CREATE TYPE payment_type AS ENUM ('upi','bank');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE payment_methods (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id         VARCHAR(255) NOT NULL,        -- stores external user ID
    type            payment_type NOT NULL,        -- 'upi' or 'bank'
    identifier      VARCHAR(255) NOT NULL,        -- UPI ID or bank account number
    bank_name       VARCHAR(150) NULL,            -- only if type = 'bank'
    ifsc_code       VARCHAR(20) NULL,             -- only if type = 'bank'
    is_default      BOOLEAN DEFAULT FALSE,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- removed FOREIGN KEY, since user_id is different type
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payment_methods;
DROP TYPE IF EXISTS payment_type;
-- +goose StatementEnd
