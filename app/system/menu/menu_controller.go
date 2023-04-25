package menu

import (
	"fiber-sqlx-arco/pkg/common/response"
	"fiber-sqlx-arco/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service Service
}

func NewController() *Controller {
	return &Controller{
		service: NewService(),
	}
}

// FindTree 查询导航菜单树
func (c *Controller) FindTree(ctx *fiber.Ctx) error {
	staffID := ctx.Locals("id").(string)
	entities, err := c.service.FindTreeByStaffID(staffID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.JSON(response.OK(entities))
}

// FindTreeAll 查询菜单树
func (c *Controller) FindTreeAll(ctx *fiber.Ctx) error {
	where := &WhereParams{
		PageSize: 10,
		Current:  1,
	}
	if err := ctx.QueryParser(where); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := utils.ValidateStruct(where); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	entities, err := c.service.FindTreeAll(where)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.JSON(response.OK(entities))
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	req := &CreateRequest{
		Sort: 100,
	}
	if err := ctx.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := utils.ValidateStruct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	if err := c.service.Create(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.JSON(response.OK(CreatedSuccess))
}

func (c *Controller) Update(ctx *fiber.Ctx) error {
	req := &UpdateRequest{}
	if err := ctx.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := utils.ValidateStruct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	if err := c.service.Update(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.JSON(response.OK(UpdatedSuccess))
}

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	deleteReq := &DeleteRequest{}
	fmt.Println(deleteReq)
	if err := ctx.BodyParser(deleteReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	fmt.Println(deleteReq)
	if err := c.service.Delete(deleteReq); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(response.OK(DeletedSuccess))
}
