package subscribers

import (
	"errors"
	"gorm.io/gorm"
	"se-school-case/pkg/model"
)

type SubscriberRepositoryInterface interface {
	GetAll() ([]model.User, error)
	AddUserEmail(user model.User) error
	CheckIfUserExists(email string) (bool, error)
}

type SubscriberRepository struct {
	db *gorm.DB
}

func NewSubscriberRepository(db *gorm.DB) SubscriberRepository {
	return SubscriberRepository{db: db}
}

func (r *SubscriberRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *SubscriberRepository) AddUserEmail(user model.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *SubscriberRepository) CheckIfUserExists(email string) (bool, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
