package service

import (
	"demorestapi/internal/entity"

	"github.com/google/uuid"
)

type UserGetter interface {
	GetUser(id string) *entity.User
}
type UserSetter interface {
	AddUser(u *entity.User) error
	UpdateUser(u *entity.User) error
}

type Service struct {
	ug UserGetter
	us UserSetter
}

func NewService(ug UserGetter, us UserSetter) *Service {
	return &Service{
		ug: ug,
		us: us,
	}
}

func (s *Service) GetUser(id string) (*entity.User, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, err
	}
	return s.ug.GetUser(id), nil
}

func (s *Service) AddUser(u *entity.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	return s.us.AddUser(u)
}

func (s *Service) UpdateUser(u *entity.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	return s.us.UpdateUser(u)
}
