package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const (
	defaultSalt = "admin_salt_2024"
)

// MD5 calculates the MD5 hash of the input string
func MD5(input string) string {
	if input == "" {
		return ""
	}
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

// MD5WithSalt calculates the MD5 hash of the input string with default salt
func MD5WithSalt(input string) string {
	return MD5WithCustomSalt(input, defaultSalt)
}

// MD5WithCustomSalt calculates the MD5 hash of the input string with custom salt
func MD5WithCustomSalt(input string, salt string) string {
	if input == "" {
		return ""
	}
	if salt == "" {
		salt = defaultSalt
	}
	return MD5(input + salt)
}

// GenerateSaltWithLength generates a random salt of specified length
func GenerateSaltWithLength(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf("failed to generate salt: %v", err)
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// GenerateSalt generates a random salt of default length (16 bytes)
func GenerateSalt() (string, error) {
	return GenerateSaltWithLength(16)
}

// Verify checks if the password matches the hashed password using default salt
func Verify(password, hashedPassword string) bool {
	if password == "" || hashedPassword == "" {
		return false
	}
	encrypted := MD5WithSalt(password)
	return encrypted == hashedPassword
}

// VerifyWithSalt checks if the password matches the hashed password using custom salt
func VerifyWithSalt(password, salt, hashedPassword string) bool {
	if password == "" || hashedPassword == "" {
		return false
	}
	encrypted := MD5WithCustomSalt(password, salt)
	return encrypted == hashedPassword
}

// MD5Multiple performs multiple rounds of MD5 hashing
func MD5Multiple(input string, times int) string {
	if input == "" || times <= 0 {
		return input
	}
	result := input
	for i := 0; i < times; i++ {
		result = MD5(result)
	}
	return result
}

// GetFileMD5 calculates the MD5 hash of a byte array
func GetFileMD5(bytes []byte) string {
	if len(bytes) == 0 {
		return ""
	}
	hash := md5.Sum(bytes)
	return hex.EncodeToString(hash[:])
}
