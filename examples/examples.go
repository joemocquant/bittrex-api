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

	printMarkets()
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
