package stocks_exchange

import (
	"encoding/json"
)

type Price struct {
	Buy         float64 `json:"buy,string"`
	Sell        float64 `json:"sell,string"`
	MarketName  string  `json:"market_name"`
	UpdatedTime int64   `json:"updated_time"`
	ServerTime  int64   `json:"server_time"`
}

func (sdk Sdk) GetPrices() ([]Price, error) {
	request := newRequest("GET", "/prices")

	responseContent, err := sdk.client.Do(request)
	if err != nil {
		return nil, err
	}

	markets := make([]Price, 0)
	err = json.Unmarshal(responseContent, &markets)
	if err != nil {
		return nil, err
	}

	return markets, nil
}
