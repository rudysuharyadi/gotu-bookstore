package services

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"gotu-bookstore/pkg/auth/dto"
	"strings"
	"time"

	"gotu-bookstore/pkg/resfmt/base_error"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
)

// Main
func sign(secretKey string, expiration int64, additionalInfo map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{
		"exp": expiration, // Token expiration time. Example: time.Now().Add(time.Hour * 1).Unix()
	}

	for key, value := range additionalInfo {
		if key == "exp" {
			return "", base_error.New(fmt.Sprintf("Invalid key for additionalInfo: %s", key))
		}
		claims[key] = value
	}

	// Create a new JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// Decode the cleaned-up PEM data
	secretKeyBytes, err := base64.StdEncoding.DecodeString(cleanupKey(secretKey))
	if err != nil {
		return "", err
	}

	// Create *ecdsa.PrivateKey
	privateKey, err := x509.ParseECPrivateKey(secretKeyBytes)
	if err != nil {
		return "", err
	}

	// Sign the token with your secret key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verify(secretKey string, tokenString string, leeway int64) (map[string]interface{}, error) {
	// Decode the cleaned-up PEM data
	secretKeyBytes, err := base64.StdEncoding.DecodeString(cleanupKey(secretKey))
	if err != nil {
		return nil, err
	}

	// Create *ecdsa.PrivateKey
	privateKey, err := x509.ParseECPrivateKey(secretKeyBytes)
	if err != nil {
		return nil, err
	}

	// Verify the token
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return &privateKey.PublicKey, nil
	}, jwt.WithLeeway(time.Second*time.Duration(leeway)))

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	} else {
		return nil, base_error.New(fmt.Sprintf("Token is not valid: %s", tokenString))
	}
}

func cleanupKey(secretKey string) string {
	secretKey = strings.ReplaceAll(secretKey, "-----BEGIN EC PRIVATE KEY-----", "")
	secretKey = strings.ReplaceAll(secretKey, "-----END EC PRIVATE KEY-----", "")
	secretKey = strings.ReplaceAll(secretKey, "-----BEGIN PUBLIC KEY-----", "")
	secretKey = strings.ReplaceAll(secretKey, "-----END PUBLIC KEY-----", "")
	secretKey = strings.ReplaceAll(secretKey, "\n", "")
	return secretKey
}

func encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Pad the data to a multiple of the block size
	padding := aes.BlockSize - len(data)%aes.BlockSize
	paddedData := append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)

	ciphertext := make([]byte, aes.BlockSize+len(paddedData))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedData)

	return ciphertext, nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Remove padding
	paddingLength := int(ciphertext[len(ciphertext)-1])
	if paddingLength < 1 || paddingLength > aes.BlockSize {
		return nil, fmt.Errorf("invalid padding length")
	}
	plaintext := ciphertext[:len(ciphertext)-paddingLength]

	return plaintext, nil
}

// SessionDTO helper
func (s AuthService) convertIntoSessionDTO(input map[string]interface{}) (*dto.SessionDTO, error) {
	var session dto.SessionDTO
	err := mapstructure.Decode(input, &session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (s AuthService) convertSessionDTOToMap(input dto.SessionDTO) (map[string]interface{}, error) {
	var output map[string]interface{}
	err := mapstructure.Decode(input, &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s AuthService) verifySessionDTO(session dto.SessionDTO) error {
	if len(session.Id) == 0 {
		return base_error.New("Invalid session ID")
	}

	return nil
}
