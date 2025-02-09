-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE Banner_Category (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    title VARCHAR(255) NULL,  -- Added NOT NULL constraint
    status BOOLEAN DEFAULT TRUE, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),  -- Use now() for default timestamp
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()  -- Use now() for default timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Banner_Category;
-- +goose StatementEnd