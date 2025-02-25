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

type Logger interface {
	Info(msg string)
	Warn(msg string, err error)
	Err(msg string, err error)
}

type Service struct {
	ug UserGetter
	us UserSetter
	l  Logger
}

func NewService(ug UserGetter, us UserSetter, l Logger) *Service {
	return &Service{
		ug: ug,
		us: us,
		l:  l,
	}
}

func (s *Service) GetUser(id string) (*entity.User, error) {
	if _, err := uuid.Parse(id); err != nil {
		s.l.Warn("uuid parse", err)
		return nil, err
	}
	return s.ug.GetUser(id), nil
}

func (s *Service) AddUser(u *entity.User) error {
	if err := u.Validate(); err != nil {
		s.l.Warn("user validation", err)
		return err
	}

	return s.us.AddUser(u)
}

func (s *Service) UpdateUser(u *entity.User) error {
	if err := u.Validate(); err != nil {
		s.l.Warn("user validation", err)
		return err
	}

	return s.us.UpdateUser(u)
}
