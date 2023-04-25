package role

import (
	"errors"
	"fiber-sqlx-arco/app/system/role_menu"
	"fiber-sqlx-arco/app/system/staff_role"
	"fiber-sqlx-arco/pkg/global"
	"fiber-sqlx-arco/pkg/utils"
	"fiber-sqlx-arco/platform/database"
	"github.com/jmoiron/sqlx"
	"time"
)

type Service interface {
	FindAll(where *WhereParams) ([]*Role, error)
	FindPage(where *WhereParams) ([]*Role, uint64, error)
	Create(req *CreateRequest) error
	Update(req *UpdateRequest) error
	Delete(req *DeleteRequest) error
	Grant(req *role_menu.Request) error
	FindAllMenus(where *role_menu.WhereParams) ([]*role_menu.RoleMenu, error)
}

type service struct {
	roleRepo      Repository
	roleMenuRepo  role_menu.Repository
	staffRoleRepo staff_role.Repository
}

func NewService() Service {
	return &service{
		roleRepo:      NewRepository(global.DB),
		roleMenuRepo:  role_menu.NewRepository(global.DB),
		staffRoleRepo: staff_role.NewRepository(global.DB),
	}
}

func (s *service) FindAll(where *WhereParams) ([]*Role, error) {
	entities, err := s.roleRepo.FindAll(where)
	return entities, err
}

func (s *service) FindPage(where *WhereParams) ([]*Role, uint64, error) {
	entities, err := s.roleRepo.FindPage(where)
	if err != nil {
		return nil, 0, err
	}
	count, err := s.roleRepo.GetCount(where)
	return entities, count, err
}

func (s *service) Create(req *CreateRequest) error {
	now := time.Now().Unix()
	operator := "root"
	entity := &Role{
		ID:        utils.GenerateID(),
		Name:      req.Name,
		Code:      req.Code,
		Status:    req.Status,
		Sort:      req.Sort,
		Remark:    req.Remark,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: operator,
		UpdatedBy: operator,
	}
	if err := s.checkFields(entity); err != nil {
		return err
	}
	err := s.roleRepo.Create(entity)
	return err
}

func (s *service) Update(req *UpdateRequest) error {
	now := time.Now().Unix()
	operator := "root"
	entity := &Role{
		ID:        req.ID,
		Name:      req.Name,
		Code:      req.Code,
		Status:    req.Status,
		Sort:      req.Sort,
		Remark:    req.Remark,
		UpdatedAt: now,
		UpdatedBy: operator,
	}
	if err := s.checkFields(entity); err != nil {
		return err
	}
	err := s.roleRepo.Update(entity)
	return err
}

func (s *service) Delete(req *DeleteRequest) error {
	//检查员工是否存在该角色
	count, _ := s.staffRoleRepo.GetCount(&staff_role.WhereParams{RoleID: req.ID})
	if count > 0 {
		return errors.New(ErrorExistStaff)
	}
	//检查角色菜单关联表中是否存在该角色
	//count, _ = s.role_menu_repo.GetCount(&role_menu.WhereParams{RoleID: req.ID})
	//if count > 0 {
	//	return errors.New(ROLE_EXIST_MENU)
	//}
	err := s.roleRepo.Delete(req)
	return err
}

func (s *service) checkFields(checkEntity *Role) error {
	name := checkEntity.Name
	code := checkEntity.Code

	entity, err := s.roleRepo.CheckFields(checkEntity)
	if err != nil {
		return nil
	}
	if entity.Name == name {
		return errors.New(ErrorNameRepeat)
	}
	if entity.Code == code {
		return errors.New(ErrorCodeRepeat)
	}
	return nil
}

func (s *service) Grant(req *role_menu.Request) error {
	now := time.Now().Unix()
	operator := "Root"
	roleId := req.RoleID
	menuIds := req.MenuIDs

	// menu_ids 为空数组，则删除所有角色菜单关联
	if len(menuIds) == 0 {
		err := s.roleMenuRepo.Delete(&role_menu.DeleteRequest{RoleID: roleId})
		return err
	} else {
		entities := []*role_menu.RoleMenu{}
		for _, menuId := range menuIds {
			entity := &role_menu.RoleMenu{
				RoleID:    roleId,
				MenuID:    menuId,
				CreatedAt: now,
				CreatedBy: operator,
			}
			entities = append(entities, entity)
		}
		// 先删除，再添加
		err := database.WithTransaction(global.DB, func(tx *sqlx.Tx) error {
			var roleIds []string
			roleIds = append(roleIds, roleId)
			count, _ := s.roleMenuRepo.GetCountWithTx(&role_menu.WhereParams{RoleIDs: roleIds}, tx)
			if count > 0 {
				if err := s.roleMenuRepo.DeleteWithTx(&role_menu.DeleteRequest{RoleID: roleId}, tx); err != nil {
					return err
				}
				if err := s.roleMenuRepo.CreateBatchWithTx(entities, tx); err != nil {
					return err
				}
			} else {
				if err := s.roleMenuRepo.CreateBatchWithTx(entities, tx); err != nil {
					return err
				}
			}
			return nil
		})
		return err
	}
}

func (s *service) FindAllMenus(where *role_menu.WhereParams) ([]*role_menu.RoleMenu, error) {
	return s.roleMenuRepo.FindAll(where)
}
