package casbin_adapter

import (
	"fiber-sqlx-arco/app/system/menu"
	"fiber-sqlx-arco/app/system/role"
	"fiber-sqlx-arco/app/system/role_menu"
	"fiber-sqlx-arco/app/system/staff"
	"fiber-sqlx-arco/app/system/staff_role"
	"fmt"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

type adapter struct {
	staffRepository     staff.Repository
	staffRoleRepository staff_role.Repository
	roleRepository      role.Repository
	roleMenuRepository  role_menu.Repository
	menuRepository      menu.Repository
}

func NewAdapterFromDB(db *sqlx.DB) persist.Adapter {
	return &adapter{
		staffRepository:     staff.NewRepository(db),
		staffRoleRepository: staff_role.NewRepository(db),
		roleRepository:      role.NewRepository(db),
		roleMenuRepository:  role_menu.NewRepository(db),
		menuRepository:      menu.NewRepository(db),
	}
}

// LoadPolicy loads policy from database.
func (a *adapter) LoadPolicy(model model.Model) error {

	err := a.loadRolePolicy(model)
	if err != nil {
		//a.logger.Zap.Errorf("Load casbin role policy error: %s", err.Error())
		return err
	}

	err = a.loadUserPolicy(model)
	if err != nil {
		//a.logger.Zap.Errorf("Load casbin user policy error: %s", err.Error())
		return err
	}

	return nil
}

// SavePolicy saves policy to database.
func (a *adapter) SavePolicy(model model.Model) error {
	return nil
}

func (a *adapter) AddPolicy(sec string, ptype string, rule []string) error {
	//TODO implement me
	panic("implement me")
}

func (a *adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	//TODO implement me
	panic("implement me")
}

func (a *adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	//TODO implement me
	panic("implement me")
}

// load role policy (p,role_id,path,method)
func (a *adapter) loadRolePolicy(model model.Model) error {
	roles, err := a.roleRepository.FindAll(&role.WhereParams{Status: "Enable"})
	if err != nil {
		return err
	}
	if len(roles) == 0 {
		return nil
	}
	// roles => role ids
	roleIDs := lo.Map[*role.Role, string](roles, func(item *role.Role, _ int) string {
		return item.ID
	})

	roleMenus, err := a.roleMenuRepository.FindAll(&role_menu.WhereParams{RoleIDs: roleIDs})
	if err != nil {
		return err
	}
	// role_menu => menu ids
	menuIDs := lo.Map[*role_menu.RoleMenu, string](roleMenus, func(item *role_menu.RoleMenu, _ int) string {
		return item.MenuID
	})

	menus, err := a.menuRepository.FindAllByIDs(menuIDs)
	if err != nil {
		return err
	}

	// 构建 p role_id path method
	for _, roleID := range roleIDs {
		for _, roleMenu := range roleMenus {
			if roleID == roleMenu.RoleID {
				for _, m := range menus {
					if roleMenu.MenuID == m.ID {
						line := fmt.Sprintf("p,%s,%s,%s", roleID, m.Path, m.Method)
						err := persist.LoadPolicyLine(line, model)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

// load user policy (g,user_id,role_id)
func (a *adapter) loadUserPolicy(model model.Model) error {
	staffs, err := a.staffRepository.FindAll(&staff.WhereParams{Status: "Enable"})
	if err != nil {
		return err
	}
	// staffs => staff ids
	staffIDs := lo.Map[*staff.Staff, string](staffs, func(item *staff.Staff, _ int) string {
		return item.ID
	})

	staffRoles, err := a.staffRoleRepository.FindAll(&staff_role.WhereParams{StaffIDs: staffIDs})

	if err != nil {
		return err
	}

	for _, staffID := range staffIDs {
		for _, staffRole := range staffRoles {
			if staffID == staffRole.StaffID {
				line := fmt.Sprintf("g,%s,%s", staffID, staffRole.RoleID)
				err := persist.LoadPolicyLine(line, model)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
