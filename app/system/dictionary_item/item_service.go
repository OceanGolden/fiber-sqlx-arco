package dictionary_item

import (
	"errors"
	"fiber-sqlx-arco/pkg/global"
	"fiber-sqlx-arco/pkg/utils"
	"time"
)

type Service interface {
	FindPage(where *WhereParams) ([]*DictionaryItem, uint64, error)
	FindById(id string) (*DictionaryItem, error)
	Create(req *CreateRequest) error
	Update(req *UpdateRequest) error
	Delete(req *DeleteRequest) error
}

type service struct {
	dictItemRepo Repository
}

func NewService() Service {
	return &service{
		dictItemRepo: NewRepository(global.DB),
	}
}

func (s *service) FindPage(where *WhereParams) ([]*DictionaryItem, uint64, error) {
	entities, err := s.dictItemRepo.FindPage(where)
	if err != nil {
		return nil, 0, err
	}
	count, err := s.dictItemRepo.GetCount(where)
	return entities, count, err
}

func (s *service) FindById(id string) (*DictionaryItem, error) {
	entity, err := s.dictItemRepo.FindById(id)
	return entity, err
}

func (s *service) Create(req *CreateRequest) error {
	now := time.Now().Unix()
	operator := "root"
	entity := &DictionaryItem{
		ID:           utils.GenerateID(),
		Label:        req.Label,
		Value:        req.Value,
		Color:        req.Color,
		DictionaryID: req.DictionaryID,
		Status:       req.Status,
		Sort:         req.Sort,
		Remark:       req.Remark,
		CreatedAt:    now,
		UpdatedAt:    now,
		CreatedBy:    operator,
		UpdatedBy:    operator,
	}
	if err := s.checkFields(entity); err != nil {
		return err
	}
	err := s.dictItemRepo.Create(entity)

	return err
}

func (s *service) Update(req *UpdateRequest) error {
	now := time.Now().Unix()
	operator := "root"
	entity := &DictionaryItem{
		ID:        req.ID,
		Label:     req.Label,
		Value:     req.Value,
		Color:     req.Color,
		Status:    req.Status,
		Sort:      req.Sort,
		Remark:    req.Remark,
		UpdatedAt: now,
		UpdatedBy: operator,
	}
	if err := s.checkFields(entity); err != nil {
		return err
	}
	err := s.dictItemRepo.Update(entity)
	return err
}

func (s *service) Delete(req *DeleteRequest) error {
	err := s.dictItemRepo.Delete(req)
	return err
}

func (s *service) checkFields(checkEntity *DictionaryItem) error {
	label := checkEntity.Label
	value := checkEntity.Value

	entity, err := s.dictItemRepo.CheckFields(checkEntity)
	if err != nil {
		return nil
	}
	if entity.Label == label {
		return errors.New(ErrorLabelRepeat)
	}
	if entity.Value == value {
		return errors.New(ErrorValueRepeat)
	}
	return nil
}
