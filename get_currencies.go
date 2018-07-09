package stocks_exchange

import (
	"encoding/json"
)

type Currency struct {
	Currency                string  `json:"currency"`
	Active                  bool    `json:"active"`
	Precision               int64   `json:"precision"`
	ApiPrecision            int64   `json:"api_precision"`
	MinimumWithdrawalAmount float64 `json:"minimum_withdrawal_amount,string"`
	MinimumDepositAmount    float64 `json:"minimum_deposit_amount,string"`
	CalculatedBalance       float64 `json:"calculated_balance,string"`
	DepositFeeCurrency      string  `json:"deposit_fee_currency"`
	DepositFeeConst         float64 `json:"deposit_fee_const,string"`
	DepositFeePercent       float64 `json:"deposit_fee_percent"`
	WithdrawalFeeCurrency   string  `json:"withdrawal_fee_currency"`
	WithdrawalFeeConst      float64 `json:"withdrawal_fee_const,string"`
	WithdrawalFeePercent    float64 `json:"withdrawal_fee_percent"`
	CurrencyLong            string  `json:"currency_long"`
	BlockExplorerUrl        string  `json:"block_explorer_url"`
}

func (sdk Sdk) GetCurrencies() ([]Currency, error) {
	request := newRequest("GET", "/currencies")

	responseContent, err := sdk.client.Do(request)
	if err != nil {
		return nil, err
	}

	currencies := make([]Currency, 0)
	err = json.Unmarshal(responseContent, &currencies)
	if err != nil {
		return nil, err
	}

	return currencies, nil
}
