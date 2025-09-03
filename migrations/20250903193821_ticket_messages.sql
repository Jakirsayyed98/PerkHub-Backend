-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE ticket_messages (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    ticket_id UUID NOT NULL REFERENCES tickets(id) ON DELETE CASCADE,
    author_type TEXT NOT NULL CHECK (author_type IN ('user','admin')), -- who wrote it
    body TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),  -- Use now() for default timestamp
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()  -- Use now() for default timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ticket_messages;
-- +goose StatementEnd