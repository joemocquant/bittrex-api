package main

import (
	"encoding/json"
	"fmt"
	"log"
	"trading/api/bittrex/publicapi"
)

var client *publicapi.PublicClient

// go run example.go
func main() {

	client = publicapi.NewPublicClient()

	// printMarkets()
	// printCurrencies()
	// printTicker()
	// printMarketSummaries()
	// printMarketSummary()
	// printOrderBook()
	printMarketHistory()
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

func printMarketSummaries() {

	res, err := client.GetMarketSummaries()

	if err != nil {
		log.Fatal(err)
	}

	prettyPrintJson(res)
}

func printMarketSummary() {

	res, err := client.GetMarketSummary("BTC-LTC")

	if err != nil {
		log.Fatal(err)
	}

	prettyPrintJson(res)
}

func printOrderBook() {

	res, err := client.GetOrderBook("BTC-LTC", "both")

	if err != nil {
		log.Fatal(err)
	}

	prettyPrintJson(res)
}

func printMarketHistory() {

	res, err := client.GetMarketHistory("BTC-LTC")

	if err != nil {
		log.Fatal(err)
	}

	prettyPrintJson(res)
}