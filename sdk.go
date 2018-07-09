package stocks_exchange

import (
	"strconv"
	"net/url"
	"net/http"
	"io/ioutil"
	"errors"
	"strings"
)

type Sdk struct {
	client Client
}

type Client interface {
	Do(request *request) ([]byte, error)
}

type client struct {
	baseUrl   string
	apiKey    string
	apiSecret string
}

type request struct {
	method     string
	path       string
	parameters url.Values
}

func newRequest(method string, path string) *request {
	return &request{
		method:     method,
		path:       path,
		parameters: url.Values{},
	}
}

func (r *request) StringParam(key string, value *string) *request {
	if value != nil {
		r.parameters.Set(key, *value)
	}
	return r
}

func (r *request) Float64Param(key string, value *float64) *request {
	if value != nil {
		r.parameters.Set(key, strconv.FormatFloat(*value, 'f', 8, 64))
	}
	return r
}

func (r *request) Int64Param(key string, value *int64) *request {
	if value != nil {
		r.parameters.Set(key, strconv.FormatInt(*value, 10))
	}
	return r
}

func (c *client) Do(request *request) ([]byte, error) {
	r, _ := c.createRequest(request)

	client := http.Client{}
	response, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	responseBody, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 300 {
		return nil, errors.New("Error " + string(response.StatusCode) + ": " + string(responseBody))
	}
	return responseBody, nil
}

func (c *client) createRequest(request *request) (*http.Request, error) {
	if request.method == "GET" {
		return c.get(request)
	}
	return c.post(request)
}

func (c *client) get(request *request) (*http.Request, error) {
	parameters := request.parameters.Encode()
	return http.NewRequest(request.method, c.baseUrl+request.path+parameters, nil)
}

func (c *client) post(request *request) (*http.Request, error) {
	form := request.parameters.Encode()
	r, err := http.NewRequest(request.method, c.baseUrl+request.path, strings.NewReader(form))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return r, nil
}

func New(apiKey string, apiSecret string) Sdk {
	return Sdk{
		client: &client{
			baseUrl:   "https://app.stocks.exchange/api2",
			apiKey:    apiKey,
			apiSecret: apiSecret,
		},
	}
}
