-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE games_data (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    code VARCHAR(100),
    url VARCHAR(255),
    name VARCHAR(50),
    isPortrait BOOLEAN,
    description VARCHAR(255),
    gamePreviewsV VARCHAR(50),
    assets VARCHAR,
    category_id VARCHAR(50),
    colorMuted VARCHAR(25),
    colorVibrant VARCHAR(25),
    status TEXT DEFAULT '1',
    privateAllowed BOOLEAN,
    rating VARCHAR(255),
    numberOfRatings VARCHAR(25),
    gamePlays VARCHAR(25),
    hasIntegratedAds  BOOLEAN,
    width VARCHAR(25),
    height VARCHAR(25),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

COMMENT ON COLUMN games_data.status IS '0 means pending, 1 means verified, 2 means rejected';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS games_data;
-- +goose StatementEnd


func InsertGames(db *sql.DB, item *request.Category) error {
	query := `
		INSERT INTO games_data ( code, url, name, isPortrait, description, gamePreviewsV, assets,category_id,colorMuted,colorVibrant,status,privateAllowed,rating,numberOfRatings,gamePlays,hasIntegratedAds,width,height,created_at,updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := db.Exec(query, item.Name, item.Description, item.Image, "1", "0", time.Now(), time.Now())
	return err
}