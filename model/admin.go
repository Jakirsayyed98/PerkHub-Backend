package model

import (
	"PerkHub/request"
	"PerkHub/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type AdminUser struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Name      *string   `json:"name,omitempty" db:"name"`   // Nullable field
	Email     *string   `json:"email,omitempty" db:"email"` // Nullable field
	Password  string    `json:"password" db:"password"`
	Verified  bool      `json:"verified" db:"verified"` // SMALLINT maps to int in Go
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func RegisterAdmin(db *sql.DB, register *request.AdminRegister) error {
	userId, err := utils.GenerateRandomUUID(15)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO adminuser (user_id, name, email, password, verified, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) ", userId, register.Name, register.Email, register.Password, true, time.Now(), time.Now())
	return err
}

func GetAdmin(db *sql.DB, email string) (*AdminUser, error) {
	var admin AdminUser

	// Use QueryRow if expecting only a single row
	err := db.QueryRow("SELECT id, user_id, name, email, password, verified, created_at, updated_at FROM adminuser WHERE email = $1", email).
		Scan(&admin.ID, &admin.UserID, &admin.Name, &admin.Email, &admin.Password, &admin.Verified, &admin.CreatedAt, &admin.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			// Handle the case where no row is found
			return nil, fmt.Errorf("no admin found with the given email")
		}
		// Other database-related errors
		return nil, err
	}

	return &admin, nil
}
