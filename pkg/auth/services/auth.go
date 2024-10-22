package services

import (
	"encoding/hex"
	"gotu-bookstore/pkg/auth/config"
	"gotu-bookstore/pkg/auth/constants"
	"gotu-bookstore/pkg/auth/dto"
	"time"

	"gotu-bookstore/pkg/resfmt/base_error"

	"golang.org/x/crypto/bcrypt"
)

type RedisInterface interface {
	Set(key string, data interface{}) error
	SetnxWithExpiry(key string, data interface{}, time int) error
	Exists(key ...string) bool
}

type AuthService struct {
	redis RedisInterface
}

func NewAuthService(
	redis RedisInterface,
) AuthService {
	return AuthService{
		redis: redis,
	}
}

// Token with SessionDTO
func (s AuthService) GenerateTokenWithSessionDTO(config config.BaseConfig, sessionDTO dto.SessionDTO) (string, error) {
	additionalInfo, err := s.convertSessionDTOToMap(sessionDTO)
	if err != nil {
		return "", err
	}

	token, err := s.GenerateToken(config, additionalInfo)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s AuthService) VerifyTokenToSessionDTO(config config.BaseConfig, token string, leeway int64) (*dto.SessionDTO, error) {
	claims, err := s.VerifyToken(config, token, leeway)
	if err != nil {
		return nil, err
	}

	session, err := s.convertIntoSessionDTO(claims)
	if err != nil {
		return nil, err
	}

	err = s.verifySessionDTO(*session)
	if err != nil {
		return nil, err
	}

	err = s.VerifyTokenValidity(token, session.Id)
	if err != nil {
		return nil, err
	}

	return session, nil
}

// Token - Generic
func (s AuthService) VerifyToken(config config.BaseConfig, token string, leeway int64) (map[string]interface{}, error) {
	claims, err := verify(config.SecretKey, token, leeway)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func (s AuthService) GenerateToken(config config.BaseConfig, additionalInfo map[string]interface{}) (string, error) {
	token, err := sign(
		config.SecretKey,
		s.GetExpirationFromNow(config.Expiration),
		additionalInfo,
	)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Invalidate token and verify validity
func (s AuthService) VerifyTokenValidity(accessToken string, userId string) error {
	exists := s.redis.Exists(
		constants.InvalidAccessTokenPrefix+accessToken,
		constants.BlacklistedUserIdPrefix+userId)
	if exists {
		return base_error.New("Invalid access token")
	}
	return nil
}

func (s AuthService) InvalidateToken(config config.BaseConfig, token string) error {
	err := s.redis.SetnxWithExpiry(
		constants.InvalidAccessTokenPrefix+token,
		time.Now(),
		int(s.GetExpirationFromNow(config.Expiration)),
	)
	if err != nil {
		return err
	}
	return nil
}

// Blacklist
func (s AuthService) BlacklistUserId(config config.BaseConfig, userId string) error {
	err := s.redis.SetnxWithExpiry(
		constants.BlacklistedUserIdPrefix+userId,
		time.Now(),
		int(config.Expiration),
	)
	if err != nil {
		return err
	}
	return nil
}

// Password
func (s AuthService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s AuthService) VerifyPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

// Utilities
func (s AuthService) GetExpirationFromNow(expiration int64) int64 {
	return time.Now().Unix() + expiration
}

// Encryption
func (s AuthService) Encrypt(plainText []byte, config config.EncryptionConfig) ([]byte, error) {
	key, err := hex.DecodeString(config.EncryptionKey)
	if err != nil {
		return nil, err
	}

	return encrypt(plainText, key)
}

func (s AuthService) Decrypt(excyptedText []byte, config config.EncryptionConfig) ([]byte, error) {
	key, err := hex.DecodeString(config.EncryptionKey)
	if err != nil {
		return nil, err
	}

	return decrypt(excyptedText, key)
}
