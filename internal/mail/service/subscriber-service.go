package service

import "se-school-case/pkg/model"

type SubscriberInterface interface {
	Add(email string) error
	GetAll() ([]model.User, error)
}
