-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE user_notification_history (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    title VARCHAR(255) NOT NULL,   -- Name of the item
    message VARCHAR(255) NOT NULL,   -- banner id of the item
    image VARCHAR(255),            -- image or path to the item's image
    click_action VARCHAR(255) NOT NULL,   -- banner id of the item
    user_id VARCHAR(255) NOT NULL,   -- banner id of the item
    status BOOLEAN DEFAULT TRUE,       -- 0 means inactive, 1 means active
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(), -- Creation timestamp
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()  -- Last update timestamp
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_notification_history;
-- +goose StatementEnd
