package mails

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"se-school-case/pkg/domain/rates"
	"se-school-case/pkg/domain/subscribers"
	"se-school-case/pkg/model"
	"se-school-case/pkg/util/constants"
)

type sendEmailToAllTestCase struct {
	name          string
	expectedUsers []model.User
	expectedRate  model.Rate
	rateError     error
	expectedError bool
}

func TestMailService_SendEmailToAll(t *testing.T) {
	constants.GOOGLE_USERNAME = "se.school.case.2024.notification@gmail.com"
	constants.GOOGLE_PASSWORD = "password"

	tests := []sendEmailToAllTestCase{
		{
			name: "Success",
			expectedUsers: []model.User{
				{Email: "email1@gmail.com"},
				{Email: "email2@gmail.com"},
			},
			expectedRate:  model.Rate{Rate: 27.5},
			expectedError: false,
		},
		{
			name: "RateError",
			expectedUsers: []model.User{
				{Email: "email1@gmail.com"},
				{Email: "email2@gmail.com"},
			},
			rateError:     errors.New("internal server error"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockSubscriberService := subscribers.NewMockSubscriberInterface(ctrl)
			mockRateService := rates.NewMockRateInterface(ctrl)

			mockSubscriberService.EXPECT().GetAll().Return(tt.expectedUsers, nil).AnyTimes()
			mockRateService.EXPECT().GetRate().Return(tt.expectedRate, tt.rateError).AnyTimes()

			mailService := NewService(mockSubscriberService, mockRateService)

			// Act
			err := mailService.SendEmailToAll("Test Subject", "../../util/resource/email.html")

			// Assert
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
