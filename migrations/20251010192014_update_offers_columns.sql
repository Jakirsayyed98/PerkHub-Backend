-- +goose Up
-- +goose StatementBegin
ALTER TABLE offers
ALTER COLUMN start_date TYPE date USING start_date::date,
ALTER COLUMN end_date TYPE date USING end_date::date;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE offers
DROP COLUMN IF EXISTS start_date,
DROP COLUMN IF EXISTS end_date;
-- +goose StatementEnd