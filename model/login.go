package model

import (
	"PerkHub/request"
	"PerkHub/utils"
	"database/sql"
	"fmt"
	"time"
)

type ResponseOTP struct {
	Return    bool     `json:"return"`
	RequestID string   `json:"request_id"`
	Message   []string `json:"message"`
}

func NewResponseOTP() *ResponseOTP {
	return &ResponseOTP{}
}

type OtpLog struct {
	Number    string    `json:"number" db:"number"`
	OTP       string    `json:"otp" db:"otp"`
	Verified  bool      `json:"verified" db:"verified"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
type UserDetail struct {
	User_id   sql.NullString `json:"user_id" db:"user_id"`
	Name      sql.NullString `json:"name" db:"name"`
	Email     sql.NullString `json:"email" db:"email"`
	Number    sql.NullString `json:"number" db:"number"`
	Gender    sql.NullString `json:"gender" db:"gender"`
	Dob       sql.NullString `json:"dob" db:"dob"`
	FCMToken  sql.NullString `json:"fcm_token" db:"fcm_token"`
	Verified  bool           `json:"verified" db:"verified"`
	Blocked   bool           `json:"blocked" db:"blocked"`
	CreatedAt sql.NullString `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at" db:"updated_at"`
}

func NewUserDetail() UserDetail {
	return UserDetail{}
}

func InsertLoginData(db *sql.DB, number string) error {
	userId, err := utils.GenerateRandomUUID(15)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO users (number,  user_id, verified) VALUES ($1, $2,true) ON CONFLICT ( number ) DO UPDATE SET verified = true ", number, userId)
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
		_, err := db.Exec(query, true, mobileNumber)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, fmt.Errorf("invalid OTP")

}

func UserDetailByMobileNumber(db *sql.DB, mobileNumber string) (*UserDetail, error) {
	query := "SELECT user_id, name, email, number, gender, dob, fcm_token, verified, created_at, updated_at FROM users WHERE number = $1"
	user := NewUserDetail()

	err := db.QueryRow(query, mobileNumber).Scan(
		&user.User_id,
		&user.Name,
		&user.Email,
		&user.Number,
		&user.Gender,
		&user.Dob,
		&user.FCMToken,
		&user.Verified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no user found with the mobile number: %s", mobileNumber)
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return &user, nil
}

func UserDetailByUserID(db *sql.DB, user_id string) (*UserDetail, error) {
	query := "SELECT user_id, name, email, number, gender, dob, fcm_token, verified, created_at, updated_at FROM users WHERE user_id = $1"
	user := NewUserDetail()

	err := db.QueryRow(query, user_id).Scan(
		&user.User_id,
		&user.Name,
		&user.Email,
		&user.Number,
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
	query := "SELECT user_id, name, email, number, gender, dob, fcm_token, verified, blocked, created_at, updated_at FROM users"

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
			&user.Gender,
			&user.Dob,
			&user.FCMToken,
			&user.Verified,
			&user.Blocked,
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

func InsertOTPRequest(db *sql.DB, number, otp string) error {
	_, err := db.Exec("INSERT INTO otp_logs (number, otp) VALUES ($1, $2)", number, otp)
	return err
}

func GetLatestOTPByNumber(db *sql.DB, number string) (string, error) {
	var otp string
	var expiresAt time.Time
	query := `SELECT otp, expires_at FROM otp_logs WHERE number = $1 ORDER BY created_at DESC LIMIT 1`
	err := db.QueryRow(query, number).Scan(&otp, &expiresAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no OTP found for number: %s", number)
		}
		return "", err
	}

	if time.Now().After(expiresAt) {
		return "", fmt.Errorf("OTP expired for number: %s", number)
	}

	return otp, nil
}

func MarkOtpVerified(db *sql.DB, number, otp string) error {
	_, err := db.Exec("UPDATE otp_logs SET verified = true WHERE number = $1 AND otp = $2", number, otp)
	return err
}
