package model

import (
	"PerkHub/request"
	"PerkHub/utils"
	"database/sql"
	"fmt"
)

type UserDetail struct {
	User_id   sql.NullString `json:"user_id" db:"user_id"`
	Name      sql.NullString `json:"name" db:"name"`
	Email     sql.NullString `json:"email" db:"email"`
	Number    sql.NullString `json:"number" db:"number"`
	OTP       sql.NullString `json:"otp" db:"otp"`
	Gender    sql.NullString `json:"gender" db:"gender"`
	Dob       sql.NullString `json:"dob" db:"dob"`
	FCMToken  sql.NullString `json:"fcm_token" db:"fcm_token"`
	Verified  sql.NullString `json:"verified" db:"verified"`
	CreatedAt sql.NullString `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at" db:"updated_at"`
}

func NewUserDetail() UserDetail {
	return UserDetail{}
}

func InsertLoginData(db *sql.DB, number, otp string) error {
	userId, err := utils.GenerateRandomUUID(15)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO users (number, otp, user_id, verified) VALUES ($1, $2, $3, 0) ON CONFLICT ( number ) DO UPDATE SET otp = EXCLUDED.otp, verified = 0 ", number, otp, userId)
	return err
}

func UpdateUserDetail(db *sql.DB, userID string, details *request.SaveUserDetailReq) error {
	_, err := db.Exec("UPDATE users SET name = $2, email = $3, gender = $4, dob = $5 WHERE user_id = $1", userID, details.Name, details.Email, details.Gender, details.DOB)
	return err
}

func VerifyOtp(db *sql.DB, mobileNumber, otp string) (bool, error) {
	var storedOTP string

	query := `SELECT otp
		FROM users
		WHERE number = $1`
	err := db.QueryRow(query, mobileNumber).Scan(&storedOTP)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("mobile number not found")
		}
		return false, err
	}

	if storedOTP == otp {
		query := "UPDATE users SET verified = $1 WHERE number = $2"
		_, err := db.Exec(query, "1", mobileNumber)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, fmt.Errorf("invalid OTP")

}

func UserDetailByMobileNumber(db *sql.DB, mobileNumber string) (*UserDetail, error) {
	query := "SELECT user_id, name, email, number, otp, gender, dob, fcm_token, verified, created_at, updated_at FROM users WHERE number = $1"
	user := NewUserDetail()

	err := db.QueryRow(query, mobileNumber).Scan(
		&user.User_id,
		&user.Name,
		&user.Email,
		&user.Number,
		&user.OTP,
		&user.Gender,
		&user.Dob,
		&user.FCMToken,
		&user.Verified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with the mobile number: %s", mobileNumber)
		}
		return nil, err
	}
	return &user, nil
}

func UserDetailByUserID(db *sql.DB, user_id string) (*UserDetail, error) {
	query := "SELECT user_id, name, email, number, otp, gender, dob, fcm_token, verified, created_at, updated_at FROM users WHERE user_id = $1"
	user := NewUserDetail()

	err := db.QueryRow(query, user_id).Scan(
		&user.User_id,
		&user.Name,
		&user.Email,
		&user.Number,
		&user.OTP,
		&user.Gender,
		&user.Dob,
		&user.FCMToken,
		&user.Verified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with the mobile number: %s", user_id)
		}
		return nil, err
	}
	return &user, nil
}

func AllUsersDetail(db *sql.DB) ([]*UserDetail, error) {
	query := "SELECT user_id, name, email, number, otp, gender, dob, fcm_token, verified, created_at, updated_at FROM users"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*UserDetail

	for rows.Next() {
		user := NewUserDetail()

		err := rows.Scan(
			&user.User_id,
			&user.Name,
			&user.Email,
			&user.Number,
			&user.OTP,
			&user.Gender,
			&user.Dob,
			&user.FCMToken,
			&user.Verified,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
