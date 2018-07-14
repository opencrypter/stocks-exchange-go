package stocks_exchange

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSdk_GetTickers(t *testing.T) {
	method, url := "GET", "/ticker"
	mockController := gomock.NewController(t)

	t.Run("It should convert api response", func(t *testing.T) {
		mockedClient := NewMockClient(mockController)
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url)

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(expectedTickersJson(), nil)

		response, _ := sdk.GetTickers()

		assert.Equal(t, expectedTickers(), response)
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

		_, err := sdk.GetTickers()

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

		_, err := sdk.GetTickers()

		assert.Error(t, err)
	})
}

func expectedTickersJson() []byte {
	return []byte(`[
  {
    "min_order_amount": "0.00000010",
    "ask": "0.00000007",
    "bid": "0.00000005",
    "last": "0.00000007",
    "lastDayAgo": "0.00000006",
    "vol": "184278225.61840923",
    "spread": "0",
    "buy_fee_percent": "0.2",
    "sell_fee_percent": "0.2",
    "market_name": "BSM_BTC",
    "market_id": 386,
    "updated_time": 1531580597,
    "server_time": 1531580597
  },
  {
    "min_order_amount": "0.00000010",
    "ask": "0.00000639",
    "bid": "0.00000639",
    "last": "0.00000639",
    "lastDayAgo": "0.00000622",
    "vol": "176.37245696",
    "spread": "0",
    "buy_fee_percent": "0.2",
    "sell_fee_percent": "0.2",
    "market_name": "TRX_BTC",
    "market_id": 500,
    "updated_time": 1531580606,
    "server_time": 1531580606
  }
]`)
}

func expectedTickers() []Ticker {
	return []Ticker{
		{
			MinOrderAmount: 0.00000010,
			Ask:            0.00000007,
			Bid:            0.00000005,
			Last:           0.00000007,
			LastDayAgo:     0.00000006,
			Volume:         184278225.61840923,
			Spread:         0,
			BuyFeePercent:  0.2,
			SellFeePercent: 0.2,
			MarketName:     "BSM_BTC",
			MarketId:       386,
			UpdatedTime:    1531580597,
			ServerTime:     1531580597,
		},
		{
			MinOrderAmount: 0.00000010,
			Ask:            0.00000639,
			Bid:            0.00000639,
			Last:           0.00000639,
			LastDayAgo:     0.00000622,
			Volume:         176.37245696,
			Spread:         0,
			BuyFeePercent:  0.2,
			SellFeePercent: 0.2,
			MarketName:     "TRX_BTC",
			MarketId:       500,
			UpdatedTime:    1531580606,
			ServerTime:     1531580606,
		},
	}
}
