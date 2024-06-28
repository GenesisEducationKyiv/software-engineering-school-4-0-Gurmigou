package rate

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"se-school-case/pkg/util/constants"
	"testing"
)

type rateFetchTestCase struct {
	name          string
	response      string
	statusCode    int
	expectedRate  float64
	expectedError string
	rateApiUrl    string
}

func mockRateAPIServer(responseBody string, statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		_, _ = w.Write([]byte(responseBody))
	})
	return httptest.NewServer(handler)
}

func TestFetchExchangeRate(t *testing.T) {
	tests := []rateFetchTestCase{
		{
			name:          "Success",
			response:      `[{"ccy":"USD","base_ccy":"UAH","buy":"36.5686","sale":"37.0700"}]`,
			statusCode:    http.StatusOK,
			expectedRate:  37.0700,
			expectedError: "",
		},
		{
			name:          "ErrorFetching",
			rateApiUrl:    "http://invalid-url",
			expectedError: "error fetching exchange rates",
		},
		{
			name:          "ErrorUnmarshaling",
			response:      `invalid json`,
			statusCode:    http.StatusOK,
			expectedError: "error unmarshaling response",
		},
		{
			name:          "RateNotFound",
			response:      `[{"ccy":"EUR","base_ccy":"UAH","buy":"40.0000","sale":"41.0000"}]`,
			statusCode:    http.StatusOK,
			expectedError: "rates not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock server if response and statusCode are provided
			if tt.response != "" {
				server := mockRateAPIServer(tt.response, tt.statusCode)
				defer server.Close()
				constants.RATE_BANK_API_URL = server.URL
			} else {
				constants.RATE_BANK_API_URL = tt.rateApiUrl
			}

			service := NewBankRateFetchService()
			rate, err := service.Fetch()

			if tt.expectedError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedRate, rate)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			}
		})
	}
}
