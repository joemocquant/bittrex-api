package publicapi

import (
	"encoding/json"
	"fmt"
)

type OrderBook struct {
	Buy  []*Order `json:"buy"`
	Sell []*Order `json:"sell"`
}

type Order struct {
	Quantity float64 `json:"Quantity"`
	Rate     float64 `json:"Rate"`
}

// Bittrex API implementation of getorderbook endpoint
//
// Endpoint: getorderbook
// Used to get retrieve the orderbook for a given market

// Parameters
// market: required  a string literal for the market (ex: BTC-LTC)
// type: required  buy, sell or both to identify the type of orderbook to return.
// depth: optional  defaults to 20 - how deep of an order book to retrieve. Max is 50 (not working)
//
// Request:
// https://bittrex.com/api/v1.1/public/getorderbook?market=BTC-LTC&type=both&depth=50
//
// Response
//
//  {
//    "success": true,
//    "message": "",
//    "result": {
//      "buy": [
//        {
//          "Quantity": 12.37000000,
//          "Rate": 0.02525000
//        }, ...
//      ],
//      "sell" : [
//        {
//          "Quantity" : 32.55412402,
//          "Rate" : 0.02540000
//        }, ...
//      ]
//    }
//  }
func (client *Client) GetOrderBook(market, typeOrder string) (*OrderBook, error) {

	params := map[string]string{
		"market": market,
		"type":   typeOrder,
	}

	resp, err := client.do("getorderbook", params)
	if err != nil {
		return nil, fmt.Errorf("Client.do: %v", err)
	}

	res := OrderBook{}

	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return &res, nil
}
