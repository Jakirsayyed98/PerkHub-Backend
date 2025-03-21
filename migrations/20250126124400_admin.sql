-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE AdminUser (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id VARCHAR(100),  -- Unique user_id with NOT NULL
    name VARCHAR(255) NULL,  -- Added NOT NULL constraint
    email VARCHAR(255)  NULL,  -- Added NOT NULL constraint
    password VARCHAR(100) UNIQUE NOT NULL,  -- Unique mobile number with NOT NULL
    verified SMALLINT DEFAULT 0,  -- Use SMALLINT for status flags
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),  -- Use now() for default timestamp
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()  -- Use now() for default timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS AdminUser;
-- +goose StatementEnd