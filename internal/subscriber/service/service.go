package service

import (
	"se-school-case/internal/subscriber/errors"
	"se-school-case/internal/subscriber/repo"
	"se-school-case/pkg/model"
)

type SubscriberService struct {
	repository repo.SubscriberRepositoryInterface
}

func NewService(repository repo.SubscriberRepositoryInterface) SubscriberService {
	return SubscriberService{repository}
}

func (s *SubscriberService) GetAll() ([]model.User, error) {
	return s.repository.GetAll()
}

func (s *SubscriberService) Add(email string) error {
	exists, err := s.repository.CheckIfUserExists(email)
	if err != nil {
		return err
	}
	if exists {
		return errors.ErrEmailAlreadyExists
	}

	user := model.User{Email: email}
	if err := s.repository.AddUserEmail(user); err != nil {
		return err
	}

	return nil
}
