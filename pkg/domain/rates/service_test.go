package rates

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"se-school-case/pkg/util/constants"
	"testing"
)

func TestRateService_GetRate_FetchError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockFetchService := NewMockRateFetchInterface(mockCtrl)
	mockRepo := NewMockRateRepository(mockCtrl)

	mockRepo.EXPECT().Where("currency_from = ? AND currency_to = ?",
		constants.DefaultCurrentFrom, constants.DefaultCurrentTo).Return(mockRepo)
	mockRepo.EXPECT().First(gomock.Any(), gomock.Any()).Return(gorm.ErrRecordNotFound)
	mockFetchService.EXPECT().FetchExchangeRate().Return(
		0.0, errors.New("fetch error")).Times(1)

	service := NewService(mockRepo, mockFetchService)

	_, err := service.GetRate()

	assert.Error(t, err)
	assert.Equal(t, "fetch error", err.Error())
}
