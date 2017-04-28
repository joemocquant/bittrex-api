package publicapi

import (
	"encoding/json"
	"fmt"
	"time"
)

type MarketHistory []*Trade

type Trade struct {
	Id        int     `json:"Id"`
	TimeStamp int64   // Unix timestamp
	Quantity  float64 `json:"Quantity"`
	Price     float64 `json:"Price"`
	Total     float64 `json:"Total"`
	FillType  string  `json:"FillType"`
	OrderType string  `json:"OrderType"`
}

// Bittrex API implementation of getmarkethistory endpoint
//
// Endpoint: getmarkethistory
// Used to retrieve the latest trades that have occured for a specific market.

// Parameters
// market: required  a string literal for the market (ex: BTC-LTC)
// count: optional  a number between 1-50 for the number of entries to return (default = 20) (not working)
//
// Request:
// https://bittrex.com/api/v1.1/public/getmarkethistory?market=BTC-DOGE&count=4
//
// Response
//
//  {
//    "success": true,
//    "message": "",
//    "result": [
//      {
//        "Id": 4861485,
//        "TimeStamp": "2017-04-26T22:27:03.633",
//        "Quantity": 40000.02,
//        "Price": 4.2e-7,
//        "Total": 0.0168,
//        "FillType": "FILL",
//        "OrderType": "BUY"
//      }, ...
//    ]
//  }
func (client *Client) GetMarketHistory(market string) (MarketHistory, error) {

	params := map[string]string{
		"market": market,
	}

	resp, err := client.do("getmarkethistory", params)
	if err != nil {
		return nil, fmt.Errorf("Client.do: %v", err)
	}

	res := make(MarketHistory, 200)

	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return res, nil
}

func (t *Trade) UnmarshalJSON(data []byte) error {

	type alias Trade
	aux := struct {
		TimeStamp string `json:"TimeStamp"`
		*alias
	}{
		alias: (*alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("json.Unmarshal: %v", err)
	}

	if timestamp, err := time.Parse("2006-01-02T15:04:05", aux.TimeStamp); err != nil {
		return fmt.Errorf("time.Parse: %v", err)
	} else {
		t.TimeStamp = int64(timestamp.Unix())
	}

	return nil
}
