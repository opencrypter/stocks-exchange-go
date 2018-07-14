# Stocks.Exchange . Go Sdk/Wrapper
[![Build Status](https://travis-ci.org/opencrypter/stocks-exchange-go.svg?branch=master)](https://travis-ci.org/opencrypter/stocks-exchange-go)
[![codecov](https://codecov.io/gh/opencrypter/stocks-exchange-go/branch/master/graph/badge.svg)](https://codecov.io/gh/opencrypter/stocks-exchange-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

An open source sdk for [Stocks.exchange](https://www.stocks.exchange/ge) written in Golang.

## Installation
After you've configured your local Go package:
```bash
go get github.com/opencrypter/stocks-exchange-go
```

## Usage
This SDK is based on the official [stocks.exchange api docs](http://help.stocks.exchange/api-integration)

You only have to call the constructor function in order to use it:

```go
package main

import "github.com/opencrypter/stocks-exchange-go"

func main() {
    stocksExchange := stocks_exchange.New("Your-api-key", "Your secret api-key")
}
```

## Available api endpoints

### Currencies
Get all available currencies with additional info.

#### Example

```go
package main

import "github.com/opencrypter/stocks-exchange-go"

func main() {
    stocksExchange := stocks_exchange.New("Your-api-key", "Your secret api-key")
    currencies, err := stocksExchange.GetCurrencies()
}
```

### Markets
Get all available currency pairs with additional info.

#### Example

```go
package main

import "github.com/opencrypter/stocks-exchange-go"

func main() {
    stocksExchange := stocks_exchange.New("Your-api-key", "Your secret api-key")
    markets , err := stocksExchange.GetMarkets()
}
```

## Tests
All is covered 100%. You can run all tests as normally you do it:
```
go test -test.v
```

## License
MIT licensed. See the LICENSE file for details.
