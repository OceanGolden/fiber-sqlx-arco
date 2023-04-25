package dictionary

import (
	"errors"
	"fiber-sqlx-arco/app/system/dictionary_item"
	"fiber-sqlx-arco/pkg/global"
	"fiber-sqlx-arco/pkg/utils"
	"time"
)

type Service interface {
	FindPage(where *WhereParams) ([]*Dictionary, uint64, error)
	Create(req *CreateRequest) error
	Update(req *UpdateRequest) error
	Delete(req *DeleteRequest) error
	FindItemsByCode(code string) ([]*dictionary_item.DictionaryItem, error)
}

type service struct {
	dictRepo     Repository
	dictItemRepo dictionary_item.Repository
}

func NewService() Service {
	return &service{
		dictRepo:     NewRepository(global.DB),
		dictItemRepo: dictionary_item.NewRepository(global.DB),
	}
}

func (s *service) FindPage(where *WhereParams) ([]*Dictionary, uint64, error) {
	entities, err := s.dictRepo.FindPage(where)
	if err != nil {
		return nil, 0, err
	}
	count, err := s.dictRepo.GetCount(where)
	return entities, count, err
}

func (s *service) Create(req *CreateRequest) error {
	now := time.Now().Unix()
	operator := "root"
	entity := &Dictionary{
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
	err := s.dictRepo.Create(entity)
	return err
}

func (s *service) Update(req *UpdateRequest) error {
	now := time.Now().Unix()
	operator := "root"
	entity := &Dictionary{
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
	err := s.dictRepo.Update(entity)
	return err
}

func (s *service) Delete(req *DeleteRequest) error {
	// 删除前 查询 子选项是否存在
	count, _ := s.dictItemRepo.GetCount(&dictionary_item.WhereParams{DictionaryID: req.ID})
	if count > 0 {
		return errors.New(ErrorExistChildren)
	}

	err := s.dictRepo.Delete(req)
	return err
}

func (s *service) checkFields(checkEntity *Dictionary) error {
	// name 唯一  code 唯一
	name := checkEntity.Name
	code := checkEntity.Code

	entity, err := s.dictRepo.CheckFields(checkEntity)
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

func (s *service) FindItemsByCode(code string) ([]*dictionary_item.DictionaryItem, error) {
	dictionary, err := s.dictRepo.FindByCode(code)
	if err != nil {
		return nil, err
	}
	entities, err := s.dictItemRepo.FindAllByDictionaryID(dictionary.ID)
	if err != nil {
		return nil, err
	}
	return entities, nil
}
