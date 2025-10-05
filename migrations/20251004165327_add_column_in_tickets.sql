-- +goose Up
-- +goose StatementBegin
ALTER TABLE tickets
ADD COLUMN priority VARCHAR(50) DEFAULT 'medium',
ADD COLUMN category VARCHAR(100);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tickets
DROP COLUMN IF EXISTS priority,
DROP COLUMN IF EXISTS category;
-- +goose StatementEnd