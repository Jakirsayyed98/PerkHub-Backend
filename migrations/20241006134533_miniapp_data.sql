-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE miniapp_data (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    miniapp_category_id TEXT NOT NULL, -- Changed to INTEGER if applicable
    miniapp_subcategory_id TEXT NOT NULL, -- Changed to INTEGER if applicable
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(255),
    description TEXT,
    cashback_terms TEXT,
    cashback_rates TEXT,
    status BOOLEAN DEFAULT FALSE,
    url_type VARCHAR(50),
    cb_active BOOLEAN DEFAULT FALSE,
    cb_percentage TEXT,
    url VARCHAR(255),
    label VARCHAR(255),
    banner VARCHAR(255),
    logo VARCHAR(255),
    macro_publisher VARCHAR(255),
    popular  BOOLEAN DEFAULT FALSE,
    trending BOOLEAN DEFAULT FALSE,
    top_cashback BOOLEAN DEFAULT FALSE,
    about TEXT,
    howitswork TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(), -- Creation timestamp
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()  -- Last update timestamp
);

-- Add comments after the table creation
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS miniapp_data;
-- +goose StatementEnd
