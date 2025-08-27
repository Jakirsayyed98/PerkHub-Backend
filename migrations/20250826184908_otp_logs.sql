-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE OTP_Logs (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    number VARCHAR(100) NOT NULL,  -- Unique mobile number with NOT NULL
    otp VARCHAR(100),  -- 5 digit OTP
    verified BOOLEAN DEFAULT FALSE,
    expires_at TIMESTAMP WITH TIME ZONE DEFAULT now() + interval '5 minutes',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),  -- Use now() for default timestamp
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()  -- Use now() for default timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS OTP_Logs;
-- +goose StatementEnd