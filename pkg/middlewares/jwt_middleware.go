package middlewares

import (
	"errors"
	"fiber-sqlx-arco/pkg/utils"
	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v3"
	"github.com/spf13/viper"
)

// JWTProtected See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:     []byte(viper.GetString("jwt.secret-key")),
		ContextKey:     "jwt",
		ErrorHandler:   jwtError,   // token无效后执行
		SuccessHandler: jwtSuccess, // token有效后执行
	}
	return jwtMiddleware.New(config)
}

func jwtSuccess(c *fiber.Ctx) error {
	metadata, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	c.Locals("id", metadata.ID)
	return c.Next()
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return fiber.NewError(fiber.StatusUnauthorized, errors.New("令牌缺失或错误！").Error())
	}
	return fiber.NewError(fiber.StatusForbidden, err.Error())
}
