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
    status TEXT DEFAULT '1', -- 0 means inactive, 1 means active
    url_type VARCHAR(50),
    cb_active TEXT DEFAULT '0',
    cb_percentage TEXT,
    url VARCHAR(255),
    label VARCHAR(255),
    banner VARCHAR(255),
    logo VARCHAR(255),
    macro_publisher VARCHAR(255),
    popular  TEXT DEFAULT '0',
    trending TEXT DEFAULT '0',
    top_cashback TEXT DEFAULT '0',
    about TEXT,
    howitswork TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(), -- Creation timestamp
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()  -- Last update timestamp
);

-- Add comments after the table creation
COMMENT ON COLUMN miniapp_data.status IS '0 means inactive, 1 means active';
COMMENT ON COLUMN miniapp_data.cb_active IS '0 means inactive, 1 means active';
COMMENT ON COLUMN miniapp_data.popular IS '0 means inactive, 1 means active';
COMMENT ON COLUMN miniapp_data.trending IS '0 means inactive, 1 means active';
COMMENT ON COLUMN miniapp_data.top_cashback IS '0 means inactive, 1 means active';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS miniapp_data;
-- +goose StatementEnd
