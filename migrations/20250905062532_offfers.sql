-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE offers (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    offer_id varchar(255),
    store_id UUID NOT NULL,-- link to miniapp_data (UUID id)
    store_name TEXT,
    title TEXT NOT NULL,
    description TEXT,
    terms_and_condition TEXT,
    coupon_code TEXT,
    image TEXT,
    type TEXT CHECK (type IN ('coupon', 'offer')) NOT NULL,
    status BOOLEAN DEFAULT TRUE, -- true = active, false = inactive
    url TEXT,
    start_date VARCHAR(255),
    end_date  VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    CONSTRAINT fk_offers_store FOREIGN KEY (store_id)
        REFERENCES miniapp_data(id) ON DELETE CASCADE
);

-- Useful indexes
CREATE INDEX IF NOT EXISTS idx_offers_offer_id ON offers(offer_id);
CREATE INDEX IF NOT EXISTS idx_offers_store_id ON offers(store_id);
CREATE INDEX IF NOT EXISTS idx_offers_status ON offers(status);
CREATE INDEX IF NOT EXISTS idx_offers_end_date ON offers(end_date);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS offers;
-- +goose StatementEnd