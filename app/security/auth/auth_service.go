package auth

import (
	"errors"
	"fiber-sqlx-arco/app/system/staff"
	"fiber-sqlx-arco/pkg/common/constants"
	"fiber-sqlx-arco/pkg/global"
	"fiber-sqlx-arco/pkg/utils"
)

type Service interface {
	Login(req *LoginRequest) (*staff.Staff, error)
	Logout(staffID string) error
}

type service struct {
	staffRepo staff.Repository
}

func NewService() Service {
	return &service{
		staffRepo: staff.NewRepository(global.DB),
	}
}

func (s *service) Login(req *LoginRequest) (*staff.Staff, error) {
	entity, err := s.staffRepo.FindOneByUsername(req.Username)
	if err != nil {
		return nil, errors.New(ErrorUsernameOrPassword)
	}
	if entity.Status != constants.ENABLE {
		return nil, errors.New(ErrorDisableStatus)
	}
	match := utils.ComparePasswords(entity.Password, req.Password)
	if !match {
		return nil, errors.New(ErrorUsernameOrPassword)
	}
	return entity, err
}

func (s *service) Logout(staffID string) error {
	return nil
}
