package auth

import (
	"fiber-sqlx-arco/app/system/menu"
	"fiber-sqlx-arco/app/system/staff"
	"fiber-sqlx-arco/pkg/common/response"
	"fiber-sqlx-arco/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	authService  Service
	staffService staff.Service
	menuService  menu.Service
}

func NewController() *Controller {
	return &Controller{
		authService:  NewService(),
		staffService: staff.NewService(),
		menuService:  menu.NewService(),
	}
}

func (c *Controller) Login(ctx *fiber.Ctx) error {
	req := &LoginRequest{}
	if err := ctx.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := utils.ValidateStruct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	entity, err := c.authService.Login(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	tokens, err := utils.GenerateNewTokens(entity.ID, entity.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(response.OK(tokens))
}

func (c *Controller) Logout(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractTokenMetadata(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	staffID := claims.ID
	err = c.authService.Logout(staffID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.JSON(response.OK(LogoutSuccess))
}

func (c *Controller) FindInfo(ctx *fiber.Ctx) error {
	claims, err := utils.ExtractTokenMetadata(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	staffID := claims.ID
	entity, err := c.staffService.FindByStaffID(staffID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	menus, permissions, err := c.menuService.FindMenuAndPermissionByStaffID(staffID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	staffInfo := &staff.Info{
		Staff:       entity,
		Menus:       menus,
		Permissions: permissions,
	}
	return ctx.JSON(response.OK(staffInfo))
}

//func (c *Controller) RefreshToken(ctx *fiber.Ctx) error {
//	claims, err := utils.ExtractTokenMetadata(ctx)
//	if err != nil {
//		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
//	}
//	staffID := claims.ID
//	now := time.Now().Unix()
//	expiresAccessToken := claims.ExpiredAt
//	if now > expiresAccessToken {
//		// Return status 401 and unauthorized error message.
//		return fiber.NewError(fiber.StatusUnauthorized, "检查token失效时间")
//	}
//}
