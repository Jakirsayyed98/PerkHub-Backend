-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE miniapp_data (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    miniapp_category_id UUID NOT NULL,
    name TEXT NOT NULL,
    icon VARCHAR(255),
    logo VARCHAR(255),
    description TEXT,
    about TEXT,
    cashback_terms TEXT,
    is_cb_active BOOLEAN NOT NULL DEFAULT true,        -- renamed for clarity
    cb_percentage NUMERIC(10,3) NOT NULL DEFAULT 0,
    url TEXT,
    url_type TEXT CHECK (url_type IN ('internal', 'external', 'deeplink')) DEFAULT 'external',
    macro_publisher UUID,
    is_active BOOLEAN NOT NULL DEFAULT true,           -- renamed for clarity
    is_popular BOOLEAN NOT NULL DEFAULT false,        -- renamed for clarity
    is_trending BOOLEAN NOT NULL DEFAULT false,       -- renamed for clarity
    is_top_cashback BOOLEAN NOT NULL DEFAULT false,   -- renamed for clarity
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_miniapp_category_id ON miniapp_data(miniapp_category_id);
CREATE INDEX idx_macro_publisher ON miniapp_data(macro_publisher);
CREATE INDEX idx_is_cb_active ON miniapp_data(is_cb_active);

-- Add comments after the table creation
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS miniapp_data;
-- +goose StatementEnd
