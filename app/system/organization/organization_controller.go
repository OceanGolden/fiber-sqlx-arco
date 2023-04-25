package organization

import (
	"fiber-sqlx-arco/pkg/common/response"
	"fiber-sqlx-arco/pkg/utils"
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

func (c *Controller) FindTree(ctx *fiber.Ctx) error {
	where := &WhereParams{
		Current:  1,
		PageSize: 10,
	}
	if err := ctx.QueryParser(where); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := utils.ValidateStruct(where); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	entities, err := c.service.FindTree(where)
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
	if err := ctx.BodyParser(deleteReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := c.service.Delete(deleteReq); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(response.OK(DeletedSuccess))
}
