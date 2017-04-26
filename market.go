package bittrex

import (
	"encoding/json"
	"fmt"
	"time"
)

type Markets []*Market

type Market struct {
	MarketCurrency     string  `json:"MarketCurrency"`
	BaseCurrency       string  `json:"BaseCurrency"`
	MarketCurrencyLong string  `json:"MarketCurrencyLong"`
	BaseCurrencyLong   string  `json:"BaseCurrencyLong"`
	MinTradeSize       float64 `json:"MinTradeSize"`
	MarketName         string  `json:"MarketName"`
	IsActive           bool    `json:"IsActive"`
	Created            int64   // Unix timestamp
}

// Bittrex API implementation of getmarkets endpoint
//
// Endpoint: getmarkets
// Used to get the open and available trading markets at Bittrex along with other meta data.

// Parameters
// None

// Request:
// https://bittrex.com/api/v1.1/public/getmarkets
//
// Response
//
//  {
//    "success" : true,
//    "message" : "",
//    "result" : [
//      {
//        "MarketCurrency" : "LTC",
//        "BaseCurrency" : "BTC",
//        "MarketCurrencyLong" : "Litecoin",
//        "BaseCurrencyLong" : "Bitcoin",
//        "MinTradeSize" : 0.01000000,
//        "MarketName" : "BTC-LTC",
//        "IsActive" : true,
//        "Created" : "2014-02-13T00:00:00"
//      }, {
//        "MarketCurrency" : "DOGE",
//        "BaseCurrency" : "BTC",
//        "MarketCurrencyLong" : "Dogecoin",
//        "BaseCurrencyLong" : "Bitcoin",
//        "MinTradeSize" : 100.00000000,
//        "MarketName" : "BTC-DOGE",
//        "IsActive" : true,
//        "Created" : "2014-02-13T00:00:00"
//      }, ...
//    ]
//  }
func (client *PublicClient) GetMarkets() (Markets, error) {

	resp, err := client.do("getmarkets", nil)
	if err != nil {
		return nil, fmt.Errorf("Client.do: %v", err)
	}

	res := make(Markets, 0)

	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return res, nil
}

func (m *Market) UnmarshalJSON(data []byte) error {

	type alias Market
	aux := struct {
		Created string `json:"Created"`
		*alias
	}{
		alias: (*alias)(m),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("json.Unmarshal: %v", err)
	}

	if timestamp, err := time.Parse("2006-01-02T15:04:05", aux.Created); err != nil {
		return fmt.Errorf("time.Parse: %v", err)
	} else {
		m.Created = int64(timestamp.Unix())
	}

	return nil
}
