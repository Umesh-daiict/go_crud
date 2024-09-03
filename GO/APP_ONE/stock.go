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

// geting 1 month ago the vlaue
func getDailyValue(ticker string) Values {
	now_data := time.Now().AddDate(0, -1, -1)
	currentFormatedDate := now_data.Format("2006-01-02")
	log.Println("sda ->", currentFormatedDate)
	resp, err := http.Get(PolygonPath + "/v1/open-close/" + strings.ToUpper(ticker) + "/" + currentFormatedDate + "?adjusted=true&apiKey=" + ApiKey)
	log.Println(PolygonPath + "/v1/open-close/" + strings.ToUpper(ticker) + "/" + currentFormatedDate + "?adjusted=true&apiKey=" + ApiKey)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	data := Values{}
	json.Unmarshal(body, &data)
	return data
}
