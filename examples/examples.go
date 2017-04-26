package main

import (
	"encoding/json"
	"fmt"
	"log"
	"trading/bittrex"
)

var client *bittrex.PublicClient

// go run example.go
func main() {

	client = bittrex.NewPublicClient()

	// printMarkets()
	// printCurrencies()
	printTicker()
}

func prettyPrintJson(msg interface{}) {

	jsonstr, err := json.MarshalIndent(msg, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(jsonstr))
}

func printMarkets() {

	res, err := client.GetMarkets()

	if err != nil {
		log.Fatal(err)
	}

	prettyPrintJson(res)
}

func printCurrencies() {

	res, err := client.GetCurrencies()

	if err != nil {
		log.Fatal(err)
	}

	prettyPrintJson(res)
}

func printTicker() {

	res, err := client.GetTicker("BTC-LTC")

	if err != nil {
		log.Fatal(err)
	}

	prettyPrintJson(res)
}
