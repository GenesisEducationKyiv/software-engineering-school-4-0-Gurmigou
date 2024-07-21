package cron_jobs_test

import (
	"encoding/json"
	"errors"
	cron_jobs "se-school-case/internal/cron-jobs"
	"se-school-case/tests/cron-jobs/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"se-school-case/pkg/model"
)

type notifySubscribersTestCase struct {
	name            string
	expectedUsers   []model.User
	expectedRate    model.Rate
	subscriberError error
	rateError       error
	expectedError   bool
	expectedEvent   cron_jobs.Event
	rabbitMQError   error
}

func TestCronJobsService_NotifySubscribers(t *testing.T) {
	tests := []notifySubscribersTestCase{
		{
			name: "Success",
			expectedUsers: []model.User{
				{Email: "email1@gmail.com"},
				{Email: "email2@gmail.com"},
			},
			expectedRate:  model.Rate{Rate: 27.5},
			expectedError: false,
			expectedEvent: cron_jobs.Event{
				EventID:     "1",
				EventType:   "RateNotification",
				AggregateID: "rate-1",
				Timestamp:   time.Now().Format(time.RFC3339),
				Data: cron_jobs.EventData{
					ExchangeRate: 27.5,
					Subscribers:  []string{"email1@gmail.com", "email2@gmail.com"},
				},
			},
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
		{
			name:            "SubscriberError",
			subscriberError: errors.New("internal server error"),
			expectedError:   true,
		},
		{
			name: "RabbitMQError",
			expectedUsers: []model.User{
				{Email: "email1@gmail.com"},
				{Email: "email2@gmail.com"},
			},
			expectedRate:  model.Rate{Rate: 27.5},
			expectedError: true,
			rabbitMQError: errors.New("failed to publish message"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockSubscriberService := mocks.NewMockSubscriberInterface(ctrl)
			mockRateService := mocks.NewMockRateInterface(ctrl)
			mockRabbitMQService := mocks.NewMockRabbitMQInterface(ctrl)

			mockSubscriberService.EXPECT().GetAll().Return(tt.expectedUsers, tt.subscriberError).AnyTimes()
			mockRateService.EXPECT().GetRate().Return(tt.expectedRate, tt.rateError).AnyTimes()

			if tt.rabbitMQError != nil {
				mockRabbitMQService.EXPECT().Publish(gomock.Any()).Return(tt.rabbitMQError).AnyTimes()
			} else {
				mockRabbitMQService.EXPECT().Publish(gomock.Any()).DoAndReturn(func(message string) error {
					var event cron_jobs.Event
					err := json.Unmarshal([]byte(message), &event)
					assert.NoError(t, err)
					assert.Equal(t, tt.expectedEvent.Data.ExchangeRate, event.Data.ExchangeRate)
					assert.ElementsMatch(t, tt.expectedEvent.Data.Subscribers, event.Data.Subscribers)
					return nil
				}).AnyTimes()
			}

			cronJobsService := cron_jobs.NewService(
				mockRabbitMQService,
				mockSubscriberService,
				mockRateService,
			)

			// Act
			err := cronJobsService.NotifySubscribers()

			// Assert
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
