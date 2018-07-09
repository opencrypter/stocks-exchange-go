package stocks_exchange

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSdk_GetCurrencies(t *testing.T) {
	method, url := "GET", "/currencies"
	mockController := gomock.NewController(t)

	t.Run("It should convert api response", func(t *testing.T) {
		mockedClient := NewMockClient(mockController)
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url)

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(expectedCurrenciesJson(), nil)

		response, _ := sdk.GetCurrencies()

		assert.Equal(t, expectedCurrencies(), response)
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

		_, err := sdk.GetCurrencies()

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

		_, err := sdk.GetCurrencies()

		assert.Error(t, err)
	})
}

func expectedCurrenciesJson() []byte {
	return []byte(`[
  		{
    		"currency": "ETHCA",
    		"active": true,
    		"precision": 8,
            "api_precision": 8,
            "minimum_withdrawal_amount": "0.00200000",
            "minimum_deposit_amount": "0.10000000",
            "deposit_fee_currency": "ETHCA",
            "deposit_fee_const": "0.001",
            "deposit_fee_percent": 0,
            "withdrawal_fee_currency": "ETHCA",
            "withdrawal_fee_const": "0.00100000",
            "withdrawal_fee_percent": 0,
            "currency_long": "Ethcash",
            "block_explorer_url": ""
  		},
  		{
            "currency": "BTC",
            "active": true,
            "precision": 8,
            "api_precision": 8,
            "minimum_withdrawal_amount": "0.00900000",
            "minimum_deposit_amount": "0.10000000",
            "deposit_fee_currency": "BTC",
            "deposit_fee_const": "0.00000000",
            "deposit_fee_percent": 0,
            "withdrawal_fee_currency": "BTC",
            "withdrawal_fee_const": "0.00150000",
            "withdrawal_fee_percent": 0.02,
            "currency_long": "Bitcoin",
            "block_explorer_url": "https://blockchain.info/tx/"
  		}
	]`)
}

func expectedCurrencies() []Currency {
	return []Currency{
		{
			Currency:                "ETHCA",
			Active:                  true,
			Precision:               8,
			ApiPrecision:            8,
			MinimumWithdrawalAmount: 0.002,
			MinimumDepositAmount:    0.1,
			DepositFeeCurrency:      "ETHCA",
			DepositFeeConst:         0.001,
			DepositFeePercent:       0,
			WithdrawalFeeConst:      0.001,
			WithdrawalFeePercent:    0,
			WithdrawalFeeCurrency:   "ETHCA",
			CurrencyLong:            "Ethcash",
			BlockExplorerUrl:        "",
		},
		{
			Currency:                "BTC",
			Active:                  true,
			Precision:               8,
			ApiPrecision:            8,
			MinimumWithdrawalAmount: 0.009,
			MinimumDepositAmount:    0.1,
			DepositFeeCurrency:      "BTC",
			DepositFeeConst:         0,
			DepositFeePercent:       0,
			WithdrawalFeePercent:    0.02,
			WithdrawalFeeConst:      0.0015,
			WithdrawalFeeCurrency:   "BTC",
			CurrencyLong:            "Bitcoin",
			BlockExplorerUrl:        "https://blockchain.info/tx/",
		},
	}
}
