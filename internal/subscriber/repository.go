package subscriber

import (
	"errors"
	"gorm.io/gorm"
	"se-school-case/pkg/model"
)

type SubscriberRepositoryInterface interface {
	CheckIfUserExists(email string) (bool, error)
}

type SubscriberRepository struct {
	db *gorm.DB
}

func NewSubscriberRepository(db *gorm.DB) SubscriberRepository {
	return SubscriberRepository{db: db}
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
