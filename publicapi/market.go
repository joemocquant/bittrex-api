package publicapi

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
	Notice             string  `json:"Notice"`
	IsSponsored        bool    `json:"IsSponsored"`
	LogoUrl            string  `json:"LogoUrl"`
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
//    "success": true,
//    "message": "",
//    "result": [
//      {
//        "MarketCurrency": "LTC",
//        "BaseCurrency": "BTC",
//        "MarketCurrencyLong": "Litecoin",
//        "BaseCurrencyLong": "Bitcoin",
//        "MinTradeSize": 0.01000000,
//        "MarketName": "BTC-LTC",
//        "IsActive": true,
//        "Created": "2014-02-13T00:00:00",
//        "Notice": "This is a crowdfund hosted by Bittrex.  See https://bittrex.com/crowdfund/lgd for more information.",
//        "IsSponsored": false,
//        "LogoUrl": "https://i.imgur.com/R29q3dD.png"
//      }, {
//        "MarketCurrency": "LGD",
//        "BaseCurrency": "BTC",
//        "MarketCurrencyLong": "Legends",
//        "BaseCurrencyLong": "Bitcoin",
//        "MinTradeSize": 1e-8,
//        "MarketName": "BTC-LGD",
//        "IsActive": true,
//        "Created": "2017-04-18T07:37:41.3",
//        "Notice": "This is a crowdfund hosted by Bittrex.  See https://bittrex.com/crowdfund/lgd for more information.",
//        "IsSponsored": null,
//        "LogoUrl": null
//      }, ...
//    ]
//  }
func (client *Client) GetMarkets() (Markets, error) {

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
