-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Reuse withdrawal_status type for old/new_status instead of new ENUM
-- If you already created withdrawal_status ENUM in withdrawal_requests migration, just use it here
CREATE TABLE withdrawal_audit_logs (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    withdrawal_id   UUID NOT NULL,
    old_status      withdrawal_status,
    new_status      withdrawal_status,
    changed_by      BIGINT NULL,                          -- admin/user/system
    note            TEXT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (withdrawal_id) REFERENCES withdrawal_requests(id)
);

CREATE INDEX idx_audit_withdrawal ON withdrawal_audit_logs(withdrawal_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS withdrawal_audit_logs;
-- +goose StatementEnd
