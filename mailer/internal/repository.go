package internal

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"uniqueIndex;not null"`
}

var CurrentRate float64 = 0

type Repository interface {
	Create(subscriber *User) error
	FindByEmail(email string) (*User, error)
	FindAll() ([]User, error)
	DeleteOne(subscriber *User) error
	DeleteAll() error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(subscriber *User) error {
	if err := r.db.Create(subscriber).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) FindByEmail(email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) DeleteOne(subscriber *User) error {
	if err := r.db.Delete(subscriber).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteAll() error {
	if err := r.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}
