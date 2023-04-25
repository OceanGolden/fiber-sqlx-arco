package organization

import (
	"errors"
	"fiber-sqlx-arco/app/system/staff"
	"fiber-sqlx-arco/pkg/common/constants"
	"fiber-sqlx-arco/pkg/global"
	"fiber-sqlx-arco/pkg/utils"
	"strings"
	"time"
)

type Service interface {
	FindTree(where *WhereParams) ([]*Organization, error)
	FindAll(where *WhereParams) ([]*Organization, error)
	Create(req *CreateRequest) error
	Update(req *UpdateRequest) error
	Delete(req *DeleteRequest) error
}

type service struct {
	organizationRepo Repository
	staffRepo        staff.Repository
}

func NewService() Service {
	return &service{
		organizationRepo: NewRepository(global.DB),
		staffRepo:        staff.NewRepository(global.DB),
	}
}

func (s *service) FindTree(Where *WhereParams) ([]*Organization, error) {
	entities, err := s.FindAll(Where)
	if err != nil {
		return nil, err
	}
	tree := s.buildTree(entities)
	return tree, err
}

func (s *service) FindAll(where *WhereParams) ([]*Organization, error) {
	// if len(where.ID) > 0 {
	// 	return s.FindAllExcludeSelfAndChildren(where)
	// }
	return s.organizationRepo.FindAll(where)
}

func (s *service) FindById(id string) (*Organization, error) {
	entity, err := s.organizationRepo.FindById(id)
	return entity, err
}

func (s *service) Create(req *CreateRequest) error {
	now := time.Now().Unix()
	operator := "root"
	entity := &Organization{
		ID:        utils.GenerateID(),
		Name:      req.Name,
		Code:      req.Code,
		ParentID:  req.ParentID,
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
	parentIds, err := s.fillParentIds(entity)
	if err != nil {
		return err
	}
	entity.ParentIDS = parentIds
	err = s.organizationRepo.Create(entity)
	return err
}

func (s *service) Update(req *UpdateRequest) error {
	now := time.Now().Unix()
	operator := "root"
	entity := &Organization{
		ID:        req.ID,
		Name:      req.Name,
		Code:      req.Code,
		ParentID:  req.ParentID,
		ParentIDS: req.ParentIDS,
		Status:    req.Status,
		Sort:      req.Sort,
		Remark:    req.Remark,
		UpdatedAt: now,
		UpdatedBy: operator,
	}
	if err := s.checkFields(entity); err != nil {
		return err
	}
	parentIds, err := s.fillParentIds(entity)
	if err != nil {
		return err
	}
	entity.ParentIDS = parentIds
	err = s.organizationRepo.Update(entity)
	return err
}

func (s *service) Delete(req *DeleteRequest) error {
	where := &WhereParams{
		ParentID: req.ID,
	}
	organizationCount, _ := s.organizationRepo.GetCount(where)
	if organizationCount > 0 {
		return errors.New(ErrorExistChildren)
	}
	staffCount, _ := s.staffRepo.GetCount(&staff.WhereParams{OrganizationID: req.ID})
	if staffCount > 0 {
		return errors.New(ErrorExistStaff)
	}
	err := s.organizationRepo.Delete(req)
	return err
}

func (s *service) fillParentIds(entity *Organization) (string, error) {
	if entity.ParentID == constants.TreeRoot || len(entity.ParentID) == 0 {
		return constants.TreeRoot, nil
	}
	parentEntity, err := s.organizationRepo.FindById(entity.ParentID)
	if err != nil {
		return "", err
	}
	parentIds := strings.Split(parentEntity.ParentIDS, ",")
	parentIds = append(parentIds, entity.ParentID)
	return strings.Join(parentIds, ","), nil
}

func (s *service) checkFields(checkEntity *Organization) error {
	name := checkEntity.Name
	code := checkEntity.Code

	if checkEntity.ID == checkEntity.ParentID {
		return errors.New(ErrorIdCantEqPid)
	}
	if checkEntity.ParentID != constants.TreeRoot {
		if checkEntity.ParentID == checkEntity.ID {
			// 父节点不能和本节点一致
			return errors.New(ErrorPidCantEqChildId)
		}

		// 如果父节点不是根节点，则父节点的父节点不能是本节点
		parentEntity, err := s.organizationRepo.FindById(checkEntity.ParentID)
		if err != nil {
			return err
		}
		if parentEntity.ParentID == checkEntity.ID {
			return errors.New(ErrorPidCantEqChildId)
		}
	}
	entity, err := s.organizationRepo.CheckFields(checkEntity)

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

func (s *service) buildTree(list []*Organization) []*Organization {
	tree := make([]*Organization, 0)
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
	//tree := lo.Filter(list, func(item *Organization, index int) bool {
	//	exist := utils.Contains(&ids, item.ParentID)
	//	return item.ParentID == constants.TreeRoot || !exist
	//})
	return s.findChildrenRecursive(tree, list)
}

func (s *service) findChildrenRecursive(tree, list []*Organization) []*Organization {
	res := make([]*Organization, 0)
	for _, node := range tree {
		children := make([]*Organization, 0)
		for _, item := range list {
			if item.ParentID == node.ID {
				children = append(children, item)
			}
		}
		//children = lo.Filter(list, func(item *Organization, index int) bool {
		//	return item.ParentID == node.ID
		//})
		if len(children) > 0 {
			node.Children = s.findChildrenRecursive(children, list)
		}
		res = append(res, node)
	}

	return res
}
