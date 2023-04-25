package utils

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"strings"
	"time"
)

const (
	ErrorExpiredTime = "签名过期，请重新登陆！"
)

// CustomClaims struct to describe metadata in JWT.
type CustomClaims struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	IssuedAt  int64  `json:"issued_at"`
	ExpiredAt int64  `json:"expired_at"`
}

func (t *CustomClaims) Valid() error {
	if time.Now().Unix() > t.ExpiredAt {
		return errors.New(ErrorExpiredTime)
	}
	return nil
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*CustomClaims, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}
	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(viper.GetString("jwt.secret-key")), nil
}
