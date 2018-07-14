package stocks_exchange

import "encoding/json"

type Ticker struct {
	MinOrderAmount float64 `json:"min_order_amount,string"`
	Ask            float64 `json:"ask,string"`
	Bid            float64 `json:"bid,string"`
	Last           float64 `json:"last,string"`
	LastDayAgo     float64 `json:"lastDayAgo,string"`
	Volume         float64 `json:"vol,string"`
	Spread         float64 `json:"spread,string"`
	BuyFeePercent  float64 `json:"buy_fee_percent,string"`
	SellFeePercent float64 `json:"sell_fee_percent,string"`
	MarketName     string  `json:"market_name"`
	MarketId       int64   `json:"market_id"`
	UpdatedTime    int64   `json:"updated_time"`
	ServerTime     int64   `json:"server_time"`
}

func (sdk Sdk) GetTickers() ([]Ticker, error) {
	request := newRequest("GET", "/ticker")

	responseContent, err := sdk.client.Do(request)
	if err != nil {
		return nil, err
	}

	markets := make([]Ticker, 0)
	err = json.Unmarshal(responseContent, &markets)
	if err != nil {
		return nil, err
	}

	return markets, nil
}
