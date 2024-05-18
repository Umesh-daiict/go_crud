package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	PolygonPath = "https://api.polygon.io"
	ApiKey      = "DvauHA_HR3_kxPO7FqYkPMusRTBw0tdA"
)

type Stock struct {
	Ticker string `json:"ticker"`
	Name   string `json:"name"`
	Price  float64
}

type Values struct {
	Open float64 `json:"open"`
}

// https://api.polygon.io/v3/reference/tickers?active=true&limit=100&apiKey=DvauHA_HR3_kxPO7FqYkPMusRTBw0tdA
func SearchStocksResult(ticker string) []Stock {
	log.Printf(PolygonPath + "/v3/reference/tickers?active=true&limit=100&apiKey=" + ApiKey + "&ticker=" + strings.ToUpper(ticker))
	resp, err := http.Get(PolygonPath + "/v3/reference/tickers?active=true&limit=100&apiKey=" + ApiKey + "&ticker=" + strings.ToUpper(ticker))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	data := struct {
		Results []Stock `json:"results"`
	}{}

	json.Unmarshal(body, &data)
	return data.Results
}

func getDailyValue(ticker string) Values {
	resp, err := http.Get(PolygonPath + "/v1/open-close/" + strings.ToUpper(ticker) + "/2024-05-15/?" + ApiKey)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	data := Values{}
	json.Unmarshal(body, &data)
	return data
}
