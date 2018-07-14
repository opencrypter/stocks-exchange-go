package stocks_exchange

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSdk_GetPrices(t *testing.T) {
	method, url := "GET", "/prices"
	mockController := gomock.NewController(t)

	t.Run("It should convert api response", func(t *testing.T) {
		mockedClient := NewMockClient(mockController)
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url)

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(expectedPricesJson(), nil)

		response, _ := sdk.GetPrices()

		assert.Equal(t, expectedPrices(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		mockedClient := NewMockClient(mockController)
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url)

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(nil, errors.New("error"))

		_, err := sdk.GetPrices()

		assert.Error(t, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		mockedClient := NewMockClient(mockController)
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url)

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(invalidJson(), nil)

		_, err := sdk.GetPrices()

		assert.Error(t, err)
	})
}

func expectedPricesJson() []byte {
	return []byte(`[
  {
    "buy": "0.00000006",
    "sell": "0.00000007",
    "market_name": "BSM_BTC",
    "updated_time": 1531588873,
    "server_time": 1531588873
  },
  {
    "buy": "0.00000037",
    "sell": "0.00000038",
    "market_name": "DOGE_BTC",
    "updated_time": 1531588873,
    "server_time": 1531588873
  }
	]`)
}

func expectedPrices() []Price {
	return []Price{
		{
			Buy:         0.00000006,
			Sell:        0.00000007,
			MarketName:  "BSM_BTC",
			UpdatedTime: 1531588873,
			ServerTime:  1531588873,
		},
		{
			Buy:         0.00000037,
			Sell:        0.00000038,
			MarketName:  "DOGE_BTC",
			UpdatedTime: 1531588873,
			ServerTime:  1531588873,
		},
	}
}
