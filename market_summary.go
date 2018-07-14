package stocks_exchange

import "encoding/json"

type getMarketSummaryQuery struct {
	currency *string
	partner  *string
}

func NewGetMarketSummaryQuery(currency string, partner string) *getMarketSummaryQuery {
	return &getMarketSummaryQuery{
		currency: &currency,
		partner:  &partner,
	}
}

func (sdk Sdk) GetMarketSummary(query *getMarketSummaryQuery) (*Market, error) {
	request := newRequest("GET", "/market_summary/"+*query.currency+"/"+*query.partner)

	responseContent, err := sdk.client.Do(request)
	if err != nil {
		return nil, err
	}

	market := &Market{}
	err = json.Unmarshal(responseContent, &market)
	if err != nil {
		return nil, err
	}

	return market, nil
}
