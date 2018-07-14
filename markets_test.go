package stocks_exchange

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSdk_GetMarkets(t *testing.T) {
	method, url := "GET", "/markets"
	mockController := gomock.NewController(t)

	t.Run("It should convert api response", func(t *testing.T) {
		mockedClient := NewMockClient(mockController)
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url)

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(expectedMarketsJson(), nil)

		response, _ := sdk.GetMarkets()

		assert.Equal(t, expectedMarkets(), response)
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

		_, err := sdk.GetMarkets()

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

		_, err := sdk.GetMarkets()

		assert.Error(t, err)
	})
}

func expectedMarketsJson() []byte {
	return []byte(`[
  {
    "currency": "BSM",
    "partner": "BTC",
    "currency_long": "Bitsum",
    "partner_long": "Bitcoin",
    "min_order_amount": "0.00000010",
    "min_buy_price": "0.00000001",
    "min_sell_price": "0.00000001",
    "buy_fee_percent": "0.2",
    "sell_fee_percent": "0.2",
    "active": true,
    "currency_precision": 8,
    "partner_precision": 8,
    "market_name": "BSM_BTC"
  },
  {
    "currency": "BTC",
    "partner": "USDT",
    "currency_long": "Bitcoin",
    "partner_long": "TetherUSD",
    "min_order_amount": "0.00000010",
    "min_buy_price": "0.00000001",
    "min_sell_price": "0.00000001",
    "buy_fee_percent": "0.2",
    "sell_fee_percent": "0.2",
    "active": true,
    "currency_precision": 8,
    "partner_precision": 8,
    "market_name": "BTC_USDT"
  }
]`)
}

func expectedMarkets() []Market {
	return []Market{
		{
			Currency:          "BSM",
			Partner:           "BTC",
			CurrencyLong:      "Bitsum",
			PartnerLong:       "Bitcoin",
			MinOrderAmount:    0.00000010,
			MinBuyPrice:       0.00000001,
			MinSellPrice:      0.00000001,
			BuyFeePercent:     0.2,
			SellFeePercent:    0.2,
			Active:            true,
			CurrencyPrecision: 8,
			PartnerPrecision:  8,
			MarketName:        "BSM_BTC",
		},
		{
			Currency:          "BTC",
			Partner:           "USDT",
			CurrencyLong:      "Bitcoin",
			PartnerLong:       "TetherUSD",
			MinOrderAmount:    0.00000010,
			MinBuyPrice:       0.00000001,
			MinSellPrice:      0.00000001,
			BuyFeePercent:     0.2,
			SellFeePercent:    0.2,
			Active:            true,
			CurrencyPrecision: 8,
			PartnerPrecision:  8,
			MarketName:        "BTC_USDT",
		},
	}
}
