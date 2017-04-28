package publicapi

import (
	"encoding/json"
	"fmt"
)

type Currencies []*Currency

type Currency struct {
	Currency        string  `json:"Currency"`
	CurrencyLong    string  `json:"CurrencyLong"`
	MinConfirmation int     `json:"MinConfirmation"`
	TxFee           float64 `json:"TxFee"`
	IsActive        bool    `json:"IsActive"`
	CoinType        string  `json:"CoinType"`
	BaseAddress     string  `json:"BaseAddress"`
	Notice          string  `json:"Notice"`
}

// Bittrex API implementation of getcurrencies endpoint
//
// Endpoint: getcurrencies
// Used to get all supported currencies at Bittrex along with other meta data.

// Parameters
// None

// Request:
// https://bittrex.com/api/v1.1/public/getcurrencies
//
// Response
//
//  {
//    "success": true,
//    "message": "",
//    "result": [
//      {
//        "Currency": "BTC",
//        "CurrencyLong": "Bitcoin",
//        "MinConfirmation": 2,
//        "TxFee": 0.00020000,
//        "IsActive": true,
//        "CoinType": "BITCOIN",
//        "BaseAddress": null
//        "Notice": null
//      }, {
//        "Currency": "LGD",
//        "CurrencyLong": "Legends",
//        "MinConfirmation": 36,
//        "TxFee": 0.01,
//        "IsActive": false,
//        "CoinType": "ETH_CONTRACT",
//        "BaseAddress": "0xfbb1b73c4f0bda4f67dca266ce6ef42f520fbb98",
//        "Notice": "Disabled until the presale ends May 15th, 2017"
//      }, ...
//    ]
//  }

func (client *Client) GetCurrencies() (Currencies, error) {

	resp, err := client.do("getcurrencies", nil)
	if err != nil {
		return nil, fmt.Errorf("Client.do: %v", err)
	}

	res := make(Currencies, 0)

	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return res, nil
}
