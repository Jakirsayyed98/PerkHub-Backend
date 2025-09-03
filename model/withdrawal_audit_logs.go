package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type WithdrawalAuditLog struct {
	ID           uuid.UUID  `db:"id"`
	WithdrawalID uuid.UUID  `db:"withdrawal_id"`
	OldStatus    string     `db:"old_status"`
	NewStatus    string     `db:"new_status"`
	ChangedBy    *uuid.UUID `db:"changed_by"`
	Note         *string    `db:"note"`
	CreatedAt    time.Time  `db:"created_at"`
}

func InsertAuditLog(db *sql.DB, withdrawalID uuid.UUID, oldStatus, newStatus string, changedBy *uuid.UUID, note *string) error {
	query := `
		INSERT INTO withdrawal_audit_logs (withdrawal_id, old_status, new_status, changed_by, note)
		VALUES ($1, $2, $3, $4, $5);
	`

	_, err := db.Exec(query, withdrawalID, oldStatus, newStatus, changedBy, note)
	if err != nil {
		return fmt.Errorf("failed to insert audit log for withdrawal %s: %w", withdrawalID, err)
	}

	return nil
}
