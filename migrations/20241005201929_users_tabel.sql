-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id VARCHAR(100),  -- Unique user_id with NOT NULL
    number VARCHAR(100) UNIQUE NOT NULL,  -- Unique mobile number with NOT NULL
    name VARCHAR(255) DEFAULT NULL,  -- Added NOT NULL constraint
    email VARCHAR(255)  DEFAULT NULL,  -- Added NOT NULL constraint
    gender VARCHAR(100) DEFAULT NULL,  -- Gender
    dob VARCHAR(100) DEFAULT NULL,  -- Changed to DATE for better type handling
    fcm_token VARCHAR(500) DEFAULT NULL,  -- Use snake_case for column names
    verified BOOLEAN DEFAULT FALSE,
    blocked BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),  -- Use now() for default timestamp
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()  -- Use now() for default timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd