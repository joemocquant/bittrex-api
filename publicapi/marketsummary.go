package publicapi

import (
	"encoding/json"
	"fmt"
	"time"
)

type MarketSummaries []*MarketSummary

type MarketSummary struct {
	MarketName     string  `json:"MarketName"`
	High           float64 `json:"High"`
	Low            float64 `json:"Low"`
	Volume         float64 `json:"Volume"`
	Last           float64 `json:"Last"`
	BaseVolume     float64 `json:"BaseVolume"`
	TimeStamp      int64   // Unix timestamp
	Bid            float64 `json:"Bid"`
	Ask            float64 `json:"Ask"`
	OpenBuyOrders  int     `json:"OpenBuyOrders"`
	OpenSellOrders int     `json:"OpenSellOrders"`
	PrevDay        float64 `json:"PrevDay"`
	Created        int64   // Unix timestamp
}

// Bittrex API implementation of getmarketsummaries endpoint
//
// Endpoint: getmarketsummaries
// Used to get the last 24 hour summary of all active exchanges

// Parameters
// None
//
// Request:
// https://bittrex.com/api/v1.1/public/getmarketsummaries
//
// Response
//
//  {
//    "success" : true,
//    "message" : "",
//    "result" : [
//      {
//        "MarketName" : "BTC-888",
//        "High" : 0.00000919,
//        "Low" : 0.00000820,
//        "Volume" : 74339.61396015,
//        "Last" : 0.00000820,
//        "BaseVolume" : 0.64966963,
//        "TimeStamp" : "2014-07-09T07:19:30.15",
//        "Bid" : 0.00000820,
//        "Ask" : 0.00000831,
//        "OpenBuyOrders" : 15,
//        "OpenSellOrders" : 15,
//        "PrevDay" : 0.00000821,
//        "Created" : "2014-03-20T06:00:00"
//      }, ...
//    ]
//  }
func (client *Client) GetMarketSummaries() (MarketSummaries, error) {

	resp, err := client.do("getmarketsummaries", nil)
	if err != nil {
		return nil, fmt.Errorf("Client.do: %v", err)
	}

	res := make(MarketSummaries, 0)

	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return res, nil
}

// Bittrex API implementation of getmarketsummary endpoint
//
// Endpoint: getmarketsummary
// Used to get the last 24 hour summary of all active exchanges

// Parameters
// market: required  a string literal for the market (ex: BTC-LTC)
//
// Request:
// https://bittrex.com/api/v1.1/public/getmarketsummary?market=btc-ltc
//
// Response
//
//  {
//    "success": true,
//    "message": "",
//    "result": [
//      {
//        "MarketName": "BTC-LTC",
//        "High": 0.01350000,
//        "Low": 0.01200000,
//        "Volume": 3833.97619253,
//        "Last": 0.01349998,
//        "BaseVolume": 47.03987026,
//        "TimeStamp": "2014-07-09T07:22:16.72",
//        "Bid": 0.01271001,
//        "Ask": 0.01291100,
//        "OpenBuyOrders": 45,
//        "OpenSellOrders": 45,
//        "PrevDay": 0.01229501,
//        "Created": "2014-02-13T00:00:00"
//      }
//    ]
//  }
func (client *Client) GetMarketSummary(market string) (*MarketSummary, error) {

	params := map[string]string{
		"market": market,
	}

	resp, err := client.do("getmarketsummary", params)
	if err != nil {
		return nil, fmt.Errorf("Client.do: %v", err)
	}

	res := make(MarketSummaries, 1)

	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return res[0], nil
}

func (m *MarketSummary) UnmarshalJSON(data []byte) error {

	type alias MarketSummary
	aux := struct {
		TimeStamp string `json:"TimeStamp"`
		Created   string `json:"Created"`
		*alias
	}{
		alias: (*alias)(m),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("json.Unmarshal: %v", err)
	}

	if timestamp, err := time.Parse("2006-01-02T15:04:05", aux.TimeStamp); err != nil {
		return fmt.Errorf("time.Parse: %v", err)
	} else {
		m.TimeStamp = int64(timestamp.Unix())
	}

	if timestamp, err := time.Parse("2006-01-02T15:04:05", aux.Created); err != nil {
		return fmt.Errorf("time.Parse: %v", err)
	} else {
		m.Created = int64(timestamp.Unix())
	}

	return nil
}
