package role

import (
	"fiber-sqlx-arco/app/system/role_menu"
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

func (c *Controller) FindAll(ctx *fiber.Ctx) error {
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
	entities, err := c.service.FindAll(where)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.JSON(response.OK(entities))
}

func (c *Controller) FindPage(ctx *fiber.Ctx) error {
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
	entities, count, err := c.service.FindPage(where)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.JSON(response.Page(entities, where.Current, where.PageSize, count))
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
	req := &DeleteRequest{}
	if err := ctx.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := utils.ValidateStruct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	if err := c.service.Delete(req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(response.OK(DeletedSuccess))
}

// GrantMenu 保存 角色菜单
func (c *Controller) GrantMenu(ctx *fiber.Ctx) error {
	req := &role_menu.Request{}
	if err := ctx.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := utils.ValidateStruct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	if err := c.service.Grant(req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(response.OK(true))
}

// FindMenus 查询 角色菜单
func (c *Controller) FindMenus(ctx *fiber.Ctx) error {
	where := &role_menu.WhereParams{}
	if err := ctx.QueryParser(where); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := utils.ValidateStruct(where); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	entities, err := c.service.FindAllMenus(where)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(response.OK(entities))
}
