package rates

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"se-school-case/pkg/util/constants"
	"testing"
)

// Create mock fetch server
func mockRateAPIServer(responseBody string, statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		_, _ = w.Write([]byte(responseBody))
	})
	return httptest.NewServer(handler)
}

func TestFetchExchangeRate_Success(t *testing.T) {
	response := `[{"ccy":"USD","base_ccy":"UAH","buy":"36.5686","sale":"37.0700"}]`
	server := mockRateAPIServer(response, http.StatusOK)
	defer server.Close()

	constants.RATE_API_URL = server.URL

	service := NewRateFetchService()
	rate, err := service.FetchExchangeRate()

	assert.NoError(t, err)
	assert.Equal(t, 37.0700, rate)
}

// Test error if fetching from server with invalid url
func TestFetchExchangeRate_ErrorFetching(t *testing.T) {
	constants.RATE_API_URL = "http://invalid-url"

	service := NewRateFetchService()
	_, err := service.FetchExchangeRate()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error fetching exchange rates")
}

// Test unmarshling error if mock fetch server return invalid json
func TestFetchExchangeRate_ErrorUnmarshaling(t *testing.T) {
	response := `invalid json`
	server := mockRateAPIServer(response, http.StatusOK)
	defer server.Close()

	constants.RATE_API_URL = server.URL

	service := NewRateFetchService()
	_, err := service.FetchExchangeRate()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error unmarshaling response")
}

// Mock fetch server responds with exchange rate EUR to UAH, not USD to UAH
// (which are the default values of constants.DefaultCurrentFrom and constants.DefaultCurrentTo)
func TestFetchExchangeRate_RateNotFound(t *testing.T) {
	response := `[{"ccy":"EUR","base_ccy":"UAH","buy":"40.0000","sale":"41.0000"}]`
	server := mockRateAPIServer(response, http.StatusOK)
	defer server.Close()

	constants.RATE_API_URL = server.URL

	service := NewRateFetchService()
	_, err := service.FetchExchangeRate()

	assert.Error(t, err)
	assert.Equal(t, "rates not found", err.Error())
}
