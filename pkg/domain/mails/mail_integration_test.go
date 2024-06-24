package mails

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"se-school-case/pkg/domain/rates"
	"se-school-case/pkg/domain/subscribers"
	"se-school-case/pkg/model"
	"se-school-case/pkg/util/constants"
	"testing"
)

func TestMailService_SendEmailToAll_Success(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSubscriberService := subscribers.NewMockSubscriberInterface(ctrl)
	mockRateService := rates.NewMockRateInterface(ctrl)

	expectedUsers := []model.User{
		{Email: "email1@gmail.com"},
		{Email: "email2@gmail.com"},
	}
	expectedRate := model.Rate{Rate: 27.5}
	constants.GOOGLE_USERNAME = "se.school.case.2024.notification@gmail.com"
	constants.GOOGLE_PASSWORD = "password"

	mockSubscriberService.EXPECT().GetAll().Return(expectedUsers, nil)
	mockRateService.EXPECT().GetRate().Return(expectedRate, nil)

	mailService := NewService(mockSubscriberService, mockRateService)

	// Act
	err := mailService.SendEmailToAll("Test Subject", "../../util/resource/email.html")

	// Assert
	assert.NoError(t, err)
}

func TestMailService_SendEmailToAll_RateError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSubscriberService := subscribers.NewMockSubscriberInterface(ctrl)
	mockRateService := rates.NewMockRateInterface(ctrl)

	expectedUsers := []model.User{
		{Email: "email1@gmail.com"},
		{Email: "email2@gmail.com"},
	}
	constants.GOOGLE_USERNAME = "se.school.case.2024.notification@gmail.com"
	constants.GOOGLE_PASSWORD = "password"

	mockSubscriberService.EXPECT().GetAll().Return(expectedUsers, nil)
	mockRateService.EXPECT().GetRate().Return(model.Rate{}, errors.New("internal server error"))

	mailService := NewService(mockSubscriberService, mockRateService)

	// Act
	err := mailService.SendEmailToAll("Test Subject", "../../util/resource/email.html")

	// Assert
	assert.Error(t, err)
}
