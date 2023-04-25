package staff

import (
	"errors"
	"fiber-sqlx-arco/app/system/staff_role"
	"fiber-sqlx-arco/pkg/common/constants"
	"fiber-sqlx-arco/pkg/global"
	"fiber-sqlx-arco/pkg/utils"
	"fiber-sqlx-arco/platform/database"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"time"
)

type Service interface {
	FindPage(where *WhereParams) ([]*Staff, uint64, error)
	Create(req *CreateRequest) error
	Update(req *UpdateRequest) error
	Delete(req *DeleteRequest) error
	AssignRole(req *staff_role.Request) error
	FindByStaffID(staffID string) (*Staff, error)
}

type service struct {
	staffRepo     Repository
	staffRoleRepo staff_role.Repository
}

func NewService() Service {
	return &service{
		staffRepo:     NewRepository(global.DB),
		staffRoleRepo: staff_role.NewRepository(global.DB),
	}
}

func (s *service) FindPage(where *WhereParams) ([]*Staff, uint64, error) {
	entities, err := s.staffRepo.FindPage(where)
	if err != nil {
		return nil, 0, err
	}
	staffIDs := lo.Map[*Staff, string](entities, func(item *Staff, _ int) string { return item.ID })
	staffRoles, err := s.staffRoleRepo.FindAll(&staff_role.WhereParams{StaffIDs: staffIDs})
	if err != nil {
		return nil, 0, err
	}
	for _, entity := range entities {
		roleIDs := lo.FilterMap[*staff_role.StaffRole, string](staffRoles, func(item *staff_role.StaffRole, _ int) (string, bool) {
			if entity.ID == item.StaffID {
				return item.RoleID, true
			}
			return "", false
		})
		entity.RoleIDs = roleIDs
	}
	count, err := s.staffRepo.GetCount(where)
	return entities, count, err
}

func (s *service) Create(req *CreateRequest) error {
	now := time.Now().Unix()
	operator := "root"
	entity := &Staff{
		ID:             utils.GenerateID(),
		Username:       req.Username,
		Password:       utils.GeneratePassword(req.Password),
		Name:           req.Name,
		Email:          req.Email,
		Mobile:         req.Mobile,
		Avatar:         req.Avatar,
		Gender:         req.Gender,
		OrganizationID: req.OrganizationID,
		PositionID:     req.PositionID,
		WorkStatus:     req.WorkStatus,
		Status:         req.Status,
		Sort:           req.Sort,
		Remark:         req.Remark,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      operator,
		UpdatedBy:      operator,
	}
	if err := s.checkFields(entity); err != nil {
		return err
	}
	err := s.staffRepo.Create(entity)
	return err
}

func (s *service) Update(req *UpdateRequest) error {
	now := time.Now().Unix()
	operator := "root"
	entity := &Staff{
		ID:             req.ID,
		Username:       req.Username,
		Name:           req.Name,
		Email:          req.Email,
		Mobile:         req.Mobile,
		Avatar:         req.Avatar,
		Gender:         req.Gender,
		OrganizationID: req.OrganizationID,
		PositionID:     req.PositionID,
		WorkStatus:     req.WorkStatus,
		Status:         req.Status,
		Sort:           req.Sort,
		Remark:         req.Remark,
		UpdatedAt:      now,
		UpdatedBy:      operator,
	}
	if err := s.checkFields(entity); err != nil {
		return err
	}
	err := s.staffRepo.Update(entity)
	return err
}

func (s *service) Delete(req *DeleteRequest) error {
	if req.ID == constants.SuperAdminId {
		return errors.New(ErrorDeleteAdminRecord)
	}
	err := s.staffRepo.Delete(req)
	return err
}

func (s *service) AssignRole(req *staff_role.Request) error {
	now := time.Now().Unix()
	operator := "Root"
	staffID := req.StaffID
	roleIDs := req.RoleIDs
	if len(roleIDs) == 0 {
		err := s.staffRoleRepo.Delete(&staff_role.DeleteRequest{StaffID: staffID})
		return err
	} else {
		entities := []*staff_role.StaffRole{}
		for _, roleID := range roleIDs {
			entity := &staff_role.StaffRole{
				StaffID:   staffID,
				RoleID:    roleID,
				CreatedAt: now,
				CreatedBy: operator,
			}
			entities = append(entities, entity)
		}
		// 检查下有没有已存在关联，先删除，再添加
		err := database.WithTransaction(global.DB, func(tx *sqlx.Tx) error {
			count, _ := s.staffRoleRepo.GetCountWithTx(&staff_role.WhereParams{StaffID: staffID}, tx)
			if count > 0 {
				if err := s.staffRoleRepo.DeleteWithTx(&staff_role.DeleteRequest{StaffID: staffID}, tx); err != nil {
					return err
				}
				if err := s.staffRoleRepo.CreateBatchWithTx(entities, tx); err != nil {
					return err
				}
			} else {
				if err := s.staffRoleRepo.CreateBatchWithTx(entities, tx); err != nil {
					return err
				}
			}
			return nil
		})
		return err
	}

}

func (s *service) checkFields(checkEntity *Staff) error {
	username := checkEntity.Username
	email := checkEntity.Email
	mobile := checkEntity.Mobile
	entity, err := s.staffRepo.CheckFields(checkEntity)
	if err != nil {
		return nil
	}
	if entity.Username == username {
		return errors.New(ErrorUsernameRepeat)
	}
	if entity.Email == email {
		return errors.New(ErrorEmailRepeat)
	}
	if entity.Mobile == mobile {
		return errors.New(ErrorMobileRepeat)
	}
	return nil
}

func (s *service) FindByStaffID(staffID string) (*Staff, error) {
	entity, err := s.staffRepo.FindById(staffID)
	if err != nil {
		return nil, errors.New(ErrorNotExist)
	}
	// 查询 roleIDs
	//staffRoles, err := s.staffRoleRepo.FindAll(&staff_role.WhereParams{StaffID: staffID})
	//if err != nil {
	//	return nil, err
	//}
	//roleIDs := lo.FilterMap[*staff_role.StaffRole, string](staffRoles, func(item *staff_role.StaffRole, _ int) (string, bool) {
	//	if entity.ID == item.StaffID {
	//		return item.RoleID, true
	//	}
	//	return "", false
	//})
	//entity.RoleIDs = roleIDs
	return entity, err
}
