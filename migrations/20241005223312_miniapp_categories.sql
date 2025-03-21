-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE miniapp_categories (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,   -- Name of the item
    description TEXT,              -- Description of the item
    image VARCHAR(500),            -- URL or path to the item's image
    status BOOLEAN DEFAULT FALSE,  -- 0 means inactive, 1 means active
    homepage_visible BOOLEAN DEFAULT FALSE,  -- 0 means inactive, 1 means active
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(), -- Creation timestamp
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()  -- Last update timestamp
);

COMMENT ON COLUMN miniapp_categories.status IS '0 means inactive, 1 means active';
COMMENT ON COLUMN miniapp_categories.homepage_visible IS '0 means inactive, 1 means active';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS miniapp_categories;
-- +goose StatementEnd