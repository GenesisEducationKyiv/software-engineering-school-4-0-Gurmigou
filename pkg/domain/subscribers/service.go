package subscribers

import (
	"errors"
	"gorm.io/gorm"
	"se-school-case/pkg/model"
	app_errors "se-school-case/pkg/util/app-error"
)

type SubscriberService struct {
	repository *gorm.DB
}

func NewService(repository *gorm.DB) SubscriberService {
	return SubscriberService{repository}
}

func (s *SubscriberService) GetAll() ([]model.User, error) {
	var users []model.User
	err := s.repository.Find(&users).Error
	return users, err
}

func (s *SubscriberService) Add(email string) error {
	exists, err := s.checkIfUserExists(email)
	if err != nil {
		return err
	}
	if exists {
		return app_errors.ErrEmailAlreadyExists
	}

	if err := s.addUserEmail(email); err != nil {
		return err
	}

	return nil
}

func (s *SubscriberService) checkIfUserExists(email string) (bool, error) {
	var user model.User
	if err := s.repository.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *SubscriberService) addUserEmail(email string) error {
	user := model.User{Email: email}
	if err := s.repository.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
