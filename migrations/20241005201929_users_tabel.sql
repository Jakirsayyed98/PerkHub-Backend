-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id VARCHAR(100),  -- Unique user_id with NOT NULL
    name VARCHAR(255) NULL,  -- Added NOT NULL constraint
    email VARCHAR(255)  NULL,  -- Added NOT NULL constraint
    number VARCHAR(100) UNIQUE NOT NULL,  -- Unique mobile number with NOT NULL
    otp VARCHAR(100),  -- 5 digit OTP
    gender VARCHAR(100),  -- Gender
    dob VARCHAR(100),  -- Changed to DATE for better type handling
    fcm_token VARCHAR(500),  -- Use snake_case for column names
    verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),  -- Use now() for default timestamp
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()  -- Use now() for default timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd