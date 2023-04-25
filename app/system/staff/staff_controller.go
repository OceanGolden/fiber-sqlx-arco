package staff

import (
	"fiber-sqlx-arco/app/system/staff_role"
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

func (c *Controller) FindPage(ctx *fiber.Ctx) error {
	where := &WhereParams{
		PageSize: 10,
		Current:  1,
	}
	if err := ctx.QueryParser(where); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := utils.ValidateStruct(*where); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	entities, count, err := c.service.FindPage(where)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(response.Page(entities, where.Current, where.PageSize, count))
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	req := &CreateRequest{
		Password: "123456",
		Sort:     100,
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
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.JSON(response.OK(DeletedSuccess))
}

func (c *Controller) AssignRole(ctx *fiber.Ctx) error {
	req := &staff_role.Request{}
	if err := ctx.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := utils.ValidateStruct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err...)
	}
	if err := c.service.AssignRole(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.JSON(response.OK(AssignSuccess))
}

// 查询用户权限

// 查询用户角色

//func (c *Controller) FindInfo(ctx *fiber.Ctx) error {
//	staffID := ctx.Locals("id").(string)
//	entity, err := c.service.FindByStaffID(staffID)
//	if err != nil {
//		return fiber.NewError(fiber.StatusBadRequest, err.Error())
//	}
//	return ctx.JSON(response.OK(entity))
//}
