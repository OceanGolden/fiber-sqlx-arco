package middlewares

import (
	"fiber-sqlx-arco/pkg/adapter/casbin_adapter"
	"fiber-sqlx-arco/pkg/global"
	"github.com/casbin/casbin/v2/persist"
	casbinMiddleware "github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

func CasbinProtected() func(*fiber.Ctx) error {
	config := casbinMiddleware.Config{
		ModelFilePath: "./pkg/configs/rbac_model.conf",
		PolicyAdapter: casbinAdapter(),
		//Enforcer:      enforcer(),
		Lookup:       lookup,
		Unauthorized: unauthorized,
		Forbidden:    forbidden,
	}
	// 按路由规则请求权限
	return casbinMiddleware.New(config).RoutePermission()
}

//func casbinAdapter() *sqlxAdapter.Adapter {
//	adapter, err := sqlxAdapter.NewAdapter(global.DB, "casbin_rule")
//	if err != nil {
//		panic(err)
//	}
//	return adapter
//}

func casbinAdapter() persist.Adapter {
	return casbin_adapter.NewAdapterFromDB(global.DB)
}

//func enforcer() *casbin.Enforcer {
//	a := casbinAdapter()
//	m, err := model.NewModelFromString(modelConf)
//	if err != nil {
//		panic(err)
//	}
//	newEnforcer, err := casbin.NewEnforcer(m, a)
//	if err != nil {
//		panic(err)
//	}
//	return newEnforcer
//}

func lookup(c *fiber.Ctx) string {
	return c.Locals("id").(string)
}

func unauthorized(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusUnauthorized, "没有权限！")
}

func forbidden(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusForbidden, "禁止访问！")
}
