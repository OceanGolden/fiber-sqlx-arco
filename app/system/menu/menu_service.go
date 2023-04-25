package menu

import (
	"errors"
	"fiber-sqlx-arco/app/system/role_menu"
	"fiber-sqlx-arco/app/system/staff_role"
	"fiber-sqlx-arco/pkg/common/constants"
	"fiber-sqlx-arco/pkg/global"
	"fiber-sqlx-arco/pkg/utils"
	"github.com/samber/lo"
	"strings"
	"time"
)

type Service interface {
	FindAll(where *WhereParams) ([]*Menu, error)
	Create(req *CreateRequest) error
	Update(req *UpdateRequest) error
	Delete(req *DeleteRequest) error
	FindTreeAll(where *WhereParams) ([]*Menu, error)
	FindTreeByStaffID(ID string) ([]*Menu, error)
	FindMenuAndPermissionByStaffID(ID string) ([]*Menu, []string, error)
}

type service struct {
	menuRepo      Repository
	staffRoleRepo staff_role.Repository
	roleMenuRepo  role_menu.Repository
}

func NewService() Service {
	return &service{
		menuRepo:      NewRepository(global.DB),
		staffRoleRepo: staff_role.NewRepository(global.DB),
		roleMenuRepo:  role_menu.NewRepository(global.DB),
	}
}

func (s *service) FindAll(where *WhereParams) ([]*Menu, error) {
	// if len(where.ID) > 0 {
	// 	return s.FindAllExcludeSelfAndChildren(where)
	// }
	return s.menuRepo.FindAll(where)
}

func (s *service) Create(req *CreateRequest) error {
	now := time.Now().Unix()
	operator := "root"
	entity := &Menu{
		ID:         utils.GenerateID(),
		Name:       req.Name,
		ParentID:   req.ParentID,
		Icon:       req.Icon,
		Path:       req.Path,
		Permission: req.Permission,
		Component:  req.Component,
		Type:       req.Type,
		Method:     req.Method,
		Status:     req.Status,
		Sort:       req.Sort,
		Remark:     req.Remark,
		CreatedAt:  now,
		UpdatedAt:  now,
		CreatedBy:  operator,
		UpdatedBy:  operator,
	}

	if err := s.checkFields(entity); err != nil {
		return err
	}
	err := s.menuRepo.Create(entity)
	return err
}

func (s *service) Update(req *UpdateRequest) error {
	now := time.Now().Unix()
	operator := "root"
	entity := &Menu{
		ID:         req.ID,
		Name:       req.Name,
		ParentID:   req.ParentID,
		Icon:       req.Icon,
		Path:       req.Path,
		Permission: req.Permission,
		Component:  req.Component,
		Type:       req.Type,
		Method:     req.Method,
		Status:     req.Status,
		Sort:       req.Sort,
		Remark:     req.Remark,
		UpdatedAt:  now,
		UpdatedBy:  operator,
	}
	if err := s.checkFields(entity); err != nil {
		return err
	}
	err := s.menuRepo.Update(entity)
	return err
}

func (s *service) Delete(req *DeleteRequest) error {
	// 查看是否有子节点
	where := &WhereParams{
		ParentID: req.ID,
	}
	count, _ := s.menuRepo.GetTotal(where)
	if count > 0 {
		return errors.New(ErrorExistChildren)
	}
	err := s.menuRepo.Delete(req)
	return err
}

func (s *service) FindTreeAll(Where *WhereParams) ([]*Menu, error) {
	entities, err := s.FindAll(Where)
	if err != nil {
		return nil, err
	}
	tree := s.buildTree(entities)
	return tree, err
}

func (s *service) FindTreeByStaffID(staffID string) ([]*Menu, error) {
	staffRoles, err := s.staffRoleRepo.FindAll(&staff_role.WhereParams{StaffID: staffID})
	//staff_role_list => role_id_list
	roleIDs := lo.Map[*staff_role.StaffRole, string](staffRoles, func(item *staff_role.StaffRole, _ int) string {
		return item.RoleID
	})
	// 通过role_ids 在role_menus中查找 menu_ids
	roleMenus, err := s.roleMenuRepo.FindAll(&role_menu.WhereParams{RoleIDs: roleIDs})
	// role_menu_list => menu_id_list
	menuIDs := lo.Map[*role_menu.RoleMenu, string](roleMenus, func(item *role_menu.RoleMenu, _ int) string {
		return item.MenuID
	})
	menusWithButton, err := s.menuRepo.FindAllByIDs(menuIDs)
	// 去除button权限菜单
	filterMenus := lo.Filter[*Menu](menusWithButton, func(item *Menu, _ int) bool { return item.Type != constants.Button })
	tree := s.buildTree(filterMenus)
	return tree, err
}

func (s *service) checkFields(checkEntity *Menu) error {
	// name 唯一
	name := checkEntity.Name
	entity, err := s.menuRepo.CheckFields(checkEntity)
	if err != nil {
		return nil
	}
	if entity.Name == name {
		return errors.New(ErrorNameRepeat)
	}
	return nil
}

func (s *service) buildTree(list []*Menu) []*Menu {
	tree := make([]*Menu, 0)
	ids := make([]string, 0)
	for _, item := range list {
		ids = append(ids, item.ID)
	}
	for _, item := range list {
		exist := utils.Contains(&ids, item.ParentID)
		if item.ParentID == constants.TreeRoot || !exist {
			tree = append(tree, item)
		}
	}
	return s.findChildrenRecursive(tree, list)
}

func (s *service) findChildrenRecursive(tree, list []*Menu) []*Menu {
	res := make([]*Menu, 0)
	for _, node := range tree {
		children := make([]*Menu, 0)
		for _, item := range list {
			if item.ParentID == node.ID {
				children = append(children, item)
			}
		}
		if len(children) > 0 {
			node.Children = s.findChildrenRecursive(children, list)
		}
		res = append(res, node)
	}

	return res
}

func (s *service) FindMenuAndPermissionByStaffID(staffID string) ([]*Menu, []string, error) {
	staffRoles, err := s.staffRoleRepo.FindAll(&staff_role.WhereParams{StaffID: staffID})
	//staff_role_list => role_id_list
	roleIDs := lo.Map[*staff_role.StaffRole, string](staffRoles, func(item *staff_role.StaffRole, _ int) string {
		return item.RoleID
	})
	// 通过role_ids 在role_menus中查找 menu_ids
	roleMenus, err := s.roleMenuRepo.FindAll(&role_menu.WhereParams{RoleIDs: roleIDs})
	// role_menu_list => menu_id_list
	menuIDs := lo.Map[*role_menu.RoleMenu, string](roleMenus, func(item *role_menu.RoleMenu, _ int) string {
		return item.MenuID
	})
	menusWithButton, err := s.menuRepo.FindAllByIDs(menuIDs)
	// 去除button权限菜单
	filterMenus := lo.Filter[*Menu](menusWithButton, func(item *Menu, _ int) bool { return item.Type != constants.Button })
	// button 权限
	buttons := lo.Filter[*Menu](menusWithButton, func(item *Menu, _ int) bool { return item.Type == constants.Button })
	permissions := lo.Map(buttons, func(item *Menu, _ int) string {
		var build strings.Builder
		build.WriteString(item.Path)
		build.WriteString(":")
		build.WriteString(item.Method)
		return build.String()
	})
	tree := s.buildTree(filterMenus)
	return tree, permissions, err
}
