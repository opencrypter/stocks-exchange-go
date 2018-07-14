package stocks_exchange

import (
	"encoding/json"
)

type Market struct {
	Currency          string  `json:"currency"`
	CurrencyLong      string  `json:"currency_long"`
	Partner           string  `json:"partner"`
	PartnerLong       string  `json:"partner_long"`
	MinOrderAmount    float64 `json:"min_order_amount,string"`
	MinBuyPrice       float64 `json:"min_buy_price,string"`
	MinSellPrice      float64 `json:"min_sell_price,string"`
	BuyFeePercent     float64 `json:"buy_fee_percent,string"`
	SellFeePercent    float64 `json:"sell_fee_percent,string"`
	Active            bool    `json:"active"`
	CurrencyPrecision int64   `json:"currency_precision"`
	PartnerPrecision  int64   `json:"partner_precision"`
	MarketName        string  `json:"market_name"`
}

func (sdk Sdk) GetMarkets() ([]Market, error) {
	request := newRequest("GET", "/markets")

	responseContent, err := sdk.client.Do(request)
	if err != nil {
		return nil, err
	}

	markets := make([]Market, 0)
	err = json.Unmarshal(responseContent, &markets)
	if err != nil {
		return nil, err
	}

	return markets, nil
}
