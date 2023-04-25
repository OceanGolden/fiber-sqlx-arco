package position

import (
	"errors"
	"fiber-sqlx-arco/pkg/global"
	"fiber-sqlx-arco/pkg/utils"
	"time"
)

type Service interface {
	FindAll(where *WhereParams) ([]*Position, error)
	FindPage(where *WhereParams) ([]*Position, uint64, error)
	FindById(id string) (*Position, error)
	Create(req *CreateRequest) error
	Update(req *UpdateRequest) error
	Delete(req *DeleteRequest) error
}

type service struct {
	positionRepo Repository
}

func NewService() Service {
	return &service{
		positionRepo: NewRepository(global.DB),
	}
}

func (s *service) FindAll(where *WhereParams) ([]*Position, error) {
	entities, err := s.positionRepo.FindAll(where)
	return entities, err
}

func (s *service) FindPage(where *WhereParams) ([]*Position, uint64, error) {
	entities, err := s.positionRepo.FindPage(where)
	if err != nil {
		return nil, 0, err
	}
	count, err := s.positionRepo.GetCount(where)
	return entities, count, err
}

func (s *service) FindById(id string) (*Position, error) {
	return s.positionRepo.FindById(id)
}

func (s *service) Create(req *CreateRequest) error {
	now := time.Now().Unix()
	operator := "root"
	entity := &Position{
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
	err := s.positionRepo.Create(entity)

	return err
}

func (s *service) Update(req *UpdateRequest) error {
	now := time.Now().Unix()
	operator := "root"
	entity := &Position{
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
	err := s.positionRepo.Update(entity)
	return err
}

func (s *service) Delete(req *DeleteRequest) error {
	err := s.positionRepo.Delete(req)
	return err
}

func (s *service) checkFields(checkEntity *Position) error {
	name := checkEntity.Name
	code := checkEntity.Code

	entity, err := s.positionRepo.CheckFields(checkEntity)
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
