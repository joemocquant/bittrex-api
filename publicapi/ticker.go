package publicapi

import (
	"encoding/json"
	"fmt"
)

type Tick struct {
	Bid  float64 `json:"Bid"`
	Ask  float64 `json:"Ask"`
	Last float64 `json:"Last"`
}

// Bittrex API implementation of getticker endpoint
//
// Endpoint: getticker
// Used to get the current tick values for a market.

// Parameters
// market: required  a string literal for the market (ex: BTC-LTC)
//
// Request:
// https://bittrex.com/api/v1.1/public/getticker
//
// Response
//
//  {
//    "success": true,
//    "message": "",
//    "result": {
//      "Bid": 2.05670368,
//      "Ask": 3.35579531,
//      "Last": 3.35579531
//    }
//  }

func (client *Client) GetTicker(market string) (*Tick, error) {

	params := map[string]string{
		"market": market,
	}

	resp, err := client.do("getticker", params)
	if err != nil {
		return nil, fmt.Errorf("Client.do: %v", err)
	}

	res := Tick{}

	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return &res, nil
}
