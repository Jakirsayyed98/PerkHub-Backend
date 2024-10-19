-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE genratedsubid_data (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    miniapp_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    subid1 VARCHAR(255) NOT NULL,
    subid2 VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(), -- Creation timestamp
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()  -- Last update timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS genratedsubid_data;
-- +goose StatementEnd
