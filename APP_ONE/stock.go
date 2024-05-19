package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
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
	now_data := time.Now()
	currentFormatedDate := now_data.Format("2024-05-15")
	log.Printf(currentFormatedDate)
	resp, err := http.Get(PolygonPath + "/v1/open-close/" + strings.ToUpper(ticker) + "/" + currentFormatedDate + "?adjusted=true&apiKey=" + ApiKey)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	data := Values{}
	json.Unmarshal(body, &data)
	return data
}
