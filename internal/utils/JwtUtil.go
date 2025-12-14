package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"flux-panel-go/internal/model/entity"
	"flux-panel-go/internal/model/localErr"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// JwtUtil JWT工具类
type JwtUtil struct {
	secretKey string
}

// 常量定义
const (
	// token有效期，7天
	ExpireTime = 7 * 24 * time.Hour
	// 算法
	Algorithm = "HmacSHA256"
)

var jwtUtil *JwtUtil

// InitJwtUtil 初始化JWT工具
func InitJwtUtil(secretKey string) {
	jwtUtil = &JwtUtil{
		secretKey: secretKey,
	}
}

// Header JWT头部
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// Payload JWT负载
type Payload struct {
	Sub    string `json:"sub"`
	Iat    int64  `json:"iat"`
	Exp    int64  `json:"exp"`
	User   string `json:"user"`
	Name   string `json:"name"`
	RoleID int    `json:"role_id"`
}

// GenerateToken 生成JWT Token
func GenerateToken(user *entity.User) (string, error) {
	now := time.Now()
	expireTime := now.Add(ExpireTime)

	// Header
	header := Header{
		Alg: Algorithm,
		Typ: "JWT",
	}
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", &localErr.CommonError{
			Code:   -2,
			ErrMsg: fmt.Sprintf("JWT header marshal failed: %v", err),
		}
	}
	encodedHeader := base64.RawURLEncoding.EncodeToString(headerJSON)

	// Payload
	payload := Payload{
		Sub:    strconv.FormatInt(int64(user.Id), 10),
		Iat:    now.Unix(),
		Exp:    expireTime.Unix(),
		User:   user.User,
		Name:   user.User,
		RoleID: user.RoleId,
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", &localErr.CommonError{
			Code:   -2,
			ErrMsg: fmt.Sprintf("JWT payload marshal failed: %v", err),
		}
	}
	encodedPayload := base64.RawURLEncoding.EncodeToString(payloadJSON)

	// Signature
	signature, err := calculateSignature(encodedHeader, encodedPayload)
	if err != nil {
		return "", &localErr.CommonError{
			Code:   -2,
			ErrMsg: fmt.Sprintf("failed to calculate signature：%v", err),
		}
	}

	// Token
	token := encodedHeader + "." + encodedPayload + "." + signature
	return token, nil
}

// ValidateToken 验证JWT Token
func ValidateToken(token string) bool {
	fmt.Printf("Starting token validation for token: %s\n", token)

	if token == "" {
		return false
	}

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false
	}

	encodedHeader := parts[0]
	encodedPayload := parts[1]
	signature := parts[2]

	// 验证签名
	expectedSignature, err := calculateSignature(encodedHeader, encodedPayload)
	if err != nil {
		fmt.Printf("Signature validation failed: %v\n", err)
		return false
	}
	if expectedSignature != signature {
		fmt.Printf("Invalid signature. Expected: %s, Got: %s\n", expectedSignature, signature)
		return false
	}

	// 验证过期时间
	decodedPayload, err := base64.RawURLEncoding.DecodeString(encodedPayload)
	if err != nil {
		return false
	}

	var payload Payload
	if err := json.Unmarshal(decodedPayload, &payload); err != nil {
		return false
	}

	now := time.Now().Unix()
	fmt.Printf("Token expiration check - Current time: %d, Token expiration: %d\n", now, payload.Exp)
	isValid := payload.Exp > now
	if isValid {
		fmt.Println("Token validation successful")
	}
	return isValid
}

// GetUserIDFromToken 从JWT Token中获取用户ID
func GetUserIDFromToken(token string) (int64, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "invalid token format",
		}
	}
	encodedPayload := parts[1]
	decodedPayload, err := base64.RawURLEncoding.DecodeString(encodedPayload)
	if err != nil {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "failed to decode payload",
		}
	}

	var payload Payload
	if err := json.Unmarshal(decodedPayload, &payload); err != nil {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "failed to unmarshal payload",
		}
	}

	userID, err := strconv.ParseInt(payload.Sub, 10, 64)
	if err != nil {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "failed to parse user ID",
		}
	}

	return userID, nil
}

// GetUserIDFromRequest 从HTTP请求中获取用户ID
func GetUserIDFromRequest(authHeader string) (int, error) {
	if authHeader == "" {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "authorization header is empty",
		}
	}

	parts := strings.Split(authHeader, ".")
	if len(parts) != 3 {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "invalid token format",
		}
	}

	encodedPayload := parts[1]
	decodedPayload, err := base64.RawURLEncoding.DecodeString(encodedPayload)
	if err != nil {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "failed to decode payload",
		}
	}

	var payload Payload
	if err := json.Unmarshal(decodedPayload, &payload); err != nil {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "failed to unmarshal payload",
		}
	}

	userID, err := strconv.Atoi(payload.Sub)
	if err != nil {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "failed to parse user ID",
		}
	}

	return userID, nil
}

// GetNameFromToken 从JWT Token中获取用户名
func GetNameFromToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", &localErr.CommonError{
			Code:   -2,
			ErrMsg: "authorization header is empty",
		}
	}

	parts := strings.Split(authHeader, ".")
	if len(parts) != 3 {
		return "", &localErr.CommonError{
			Code:   -2,
			ErrMsg: "invalid token format",
		}
	}

	encodedPayload := parts[1]
	decodedPayload, err := base64.RawURLEncoding.DecodeString(encodedPayload)
	if err != nil {
		return "", &localErr.CommonError{
			Code:   -2,
			ErrMsg: "failed to decode payload",
		}
	}

	var payload Payload
	if err := json.Unmarshal(decodedPayload, &payload); err != nil {
		return "", &localErr.CommonError{
			Code:   -2,
			ErrMsg: "failed to unmarshal payload",
		}
	}

	return payload.Name, nil
}

// GetRoleIDFromToken 从JWT Token中获取角色ID
func GetRoleIDFromToken(token string) (int, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "invalid token format",
		}
	}

	encodedPayload := parts[1]
	decodedPayload, err := base64.RawURLEncoding.DecodeString(encodedPayload)
	if err != nil {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "failed to decode payload",
		}
	}

	var payload Payload
	if err := json.Unmarshal(decodedPayload, &payload); err != nil {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "failed to unmarshal payload",
		}
	}

	return payload.RoleID, nil
}

// GetRoleIDFromRequest 从HTTP请求中获取角色ID
func GetRoleIDFromRequest(authHeader string) (int, error) {
	if authHeader == "" {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "authorization header is empty",
		}
	}

	parts := strings.Split(authHeader, ".")
	if len(parts) != 3 {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "invalid token format",
		}
	}

	encodedPayload := parts[1]
	decodedPayload, err := base64.RawURLEncoding.DecodeString(encodedPayload)
	if err != nil {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "failed to decode payload",
		}
	}

	var payload Payload
	if err := json.Unmarshal(decodedPayload, &payload); err != nil {
		return 0, &localErr.CommonError{
			Code:   -2,
			ErrMsg: "failed to unmarshal payload",
		}
	}

	return payload.RoleID, nil
}

// calculateSignature 计算签名
func calculateSignature(encodedHeader, encodedPayload string) (string, error) {
	if jwtUtil == nil {
		return "", &localErr.CommonError{
			Code:   -2,
			ErrMsg: "JWT utility not initialized",
		}
	}

	content := encodedHeader + "." + encodedPayload
	h := hmac.New(sha256.New, []byte(jwtUtil.secretKey))
	h.Write([]byte(content))
	signatureBytes := h.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(signatureBytes), nil
}
