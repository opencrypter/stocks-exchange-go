package stocks_exchange

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSdk_GetMarketSummaryummary(t *testing.T) {
	method, url := "GET", "/market_summary/BSM/BTC"
	mockController := gomock.NewController(t)

	t.Run("It should convert api response", func(t *testing.T) {
		mockedClient := NewMockClient(mockController)
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url)

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(expectedMarketSummaryJson(), nil)

		query := NewGetMarketSummaryQuery("BSM", "BTC")
		response, _ := sdk.GetMarketSummary(query)

		assert.Equal(t, expectedMarketSummary(), response)
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

		query := NewGetMarketSummaryQuery("BSM", "BTC")
		_, err := sdk.GetMarketSummary(query)

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

		query := NewGetMarketSummaryQuery("BSM", "BTC")
		_, err := sdk.GetMarketSummary(query)

		assert.Error(t, err)
	})
}

func expectedMarketSummaryJson() []byte {
	return []byte(`{
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
}`)
}

func expectedMarketSummary() *Market {
	return &Market{
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
	}
}
