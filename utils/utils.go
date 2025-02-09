package utils

import (
	"PerkHub/connection"
	"PerkHub/constants"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	mathRand "math/rand"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/leebenson/conform"
)

func GenerateNumber(length int) string {
	mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	code := ""
	for i := 0; i < length; i++ {
		digit := mathRand.Intn(10) + 48
		code += string(rune(digit))
	}

	return code
}

func uint64ToBytes(num uint64) []byte {
	buf := make([]byte, 8)
	for i := uint(0); i < 8; i++ {
		buf[i] = byte(num >> (56 - (i * 8)))
	}
	return buf
}

// GenerateRandomUUID generates a random UUID (version 4) without using external libraries.
func GenerateRandomUUID(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length should be greater than 0")
	}

	uuid := make([]byte, 16)

	// Set the version (bits 12-15) to 0100 for UUID version 4.
	uuid[6] = 0x40 | (uuid[6] & 0xF)

	// Set the variant (bits 16-17) to 10 for RFC 4122.
	uuid[8] = 0x80 | (uuid[8] & 0x3F)

	// Fill the first 8 bytes with the current timestamp.
	timestampBytes := uint64ToBytes(uint64(time.Now().UnixNano() / 100))
	copy(uuid[0:8], timestampBytes[2:])

	// Fill the rest of the UUID with random bytes.
	_, err := rand.Read(uuid[8:])
	if err != nil {
		return "", err
	}

	// Convert to UUID string format.
	uuidString := fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])

	// Trim or pad the string to achieve the desired length
	if len(uuidString) >= length {
		return uuidString[:length], nil
	}

	return uuidString + fmt.Sprintf("%0*s", length-len(uuidString), ""), nil
}

type Claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

type SecretClaims struct {
	Key string `json:"key"`
	IV  string `json:"iv"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(user_id string) (string, error) {
	secretKey := []byte(constants.JWT_KEY)
	claims := jwt.MapClaims{
		"user_id": fmt.Sprintf("%s|%s", user_id, string(secretKey)),
		"iss":     "perkhub",
		"exp":     time.Now().Add(time.Hour * 10).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error signing the token:", err)
		return "", err
	}

	return signedToken, nil
}

func VerifyJwt(tokenString string, secret string) (string, error) {
	jwtKey := []byte(secret)
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "Authorization error", err
	}

	if !tkn.Valid {
		return "Authorization error", errors.New("invalid token")
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return "Authorization error", errors.New("token expired")
	}

	return claims.UserId, nil
}

func ValidateStruct(payload interface{}) (err error) {
	if err := conform.Strings(payload); err != nil {
		return err
	}

	return nil
}

func SaveFile(c *gin.Context, file *multipart.FileHeader) (string, error) {
	timestamp := time.Now().Format("20060102150405")
	sanitizedFileName := strings.ReplaceAll(file.Filename, " ", "_")
	filename := fmt.Sprintf("%s%s", timestamp, sanitizedFileName)
	fmt.Println(filename)

	desti := filepath.Join("./files", filename)
	if err := c.SaveUploadedFile(file, desti); err != nil {

		return "", err
	}

	imageURL := filename

	return imageURL, nil
}

// Encrypt encrypts plain text using AES
func Encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher([]byte(constants.JWT_KEY))
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))

	// Return as base64 encoded string
	return base64.URLEncoding.EncodeToString(cipherText), nil
}

// Decrypt decrypts AES-encrypted text
func Decrypt(encryptedText string) (string, error) {
	cipherText, err := base64.URLEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(constants.JWT_KEY))
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

func UploadFileOnServer(files []*multipart.FileHeader, awsInstance *connection.Aws) (string, error) {
	if len(files) > 0 {

		fileHeader := files[0]
		f, err := fileHeader.Open()
		if err != nil {
			return "", err
		}
		defer f.Close()

		image, err := awsInstance.UploadFile(f, fileHeader.Filename, constants.AWSBucketName, constants.AWSSecretAccessKey)
		// image, err = utils.SaveFile(c, file)
		if err != nil {
			return "", err
		}
		return image, nil
	}
	return "", nil
}

func ImageUrlGenerator(image string) string {
	if image == "" {
		return ""
	}
	return constants.AWSCloudFrontURL + constants.AWSSecretAccessKey + "/" + image
}
