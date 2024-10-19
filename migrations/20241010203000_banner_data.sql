-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE banner_data (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,   -- Name of the item
    banner_id VARCHAR(255) NOT NULL,   -- banner id of the item
    image VARCHAR(500),            -- image or path to the item's image
    url VARCHAR(255) NOT NULL,   -- url of the item
    status TEXT DEFAULT '1',       -- 0 means inactive, 1 means active
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(), -- Creation timestamp
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()  -- Last update timestamp
);

COMMENT ON COLUMN banner_data.status IS '0 means inactive, 1 means active';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS banner_data;
-- +goose StatementEnd
