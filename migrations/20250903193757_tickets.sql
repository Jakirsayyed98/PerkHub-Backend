-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE tickets (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id VARCHAR(100),  -- Unique user_id with NOT NULL
    subject TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'open', -- open/closed
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),  -- Use now() for default timestamp
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()  -- Use now() for default timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tickets;
-- +goose StatementEnd