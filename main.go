
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type jsonData []struct {
		Symbol                string  `json:"symbol"`
		CompanyName           string  `json:"companyName"`
		PrimaryExchange       string  `json:"primaryExchange"`
		Sector                string  `json:"sector"`
		CalculationPrice      string  `json:"calculationPrice"`
		Open                  float64 `json:"open"`
		OpenTime              int64   `json:"openTime"`
		Close                 float64 `json:"close"`
		CloseTime             int64   `json:"closeTime"`
		High                  float64 `json:"high"`
		Low                   float64 `json:"low"`
		LatestPrice           float64 `json:"latestPrice"`
		LatestSource          string  `json:"latestSource"`
		LatestTime            string  `json:"latestTime"`
		LatestUpdate          int64   `json:"latestUpdate"`
		LatestVolume          float64     `json:"latestVolume"`
		IexRealtimePrice      float64 `json:"iexRealtimePrice"`
		IexRealtimeSize       float64     `json:"iexRealtimeSize"`
		IexLastUpdated        int64   `json:"iexLastUpdated"`
		DelayedPrice          float64 `json:"delayedPrice"`
		DelayedPriceTime      int64   `json:"delayedPriceTime"`
		ExtendedPrice         float64 `json:"extendedPrice"`
		ExtendedChange        float64     `json:"extendedChange"`
		ExtendedChangePercent float64     `json:"extendedChangePercent"`
		ExtendedPriceTime     int64   `json:"extendedPriceTime"`
		PreviousClose         float64 `json:"previousClose"`
		Change                float64 `json:"change"`
		ChangePercent         float64 `json:"changePercent"`
		IexMarketPercent      float64 `json:"iexMarketPercent"`
		IexVolume             float64     `json:"iexVolume"`
		AvgTotalVolume        float64     `json:"avgTotalVolume"`
		IexBidPrice           float64     `json:"iexBidPrice"`
		IexBidSize            float64     `json:"iexBidSize"`
		IexAskPrice           float64     `json:"iexAskPrice"`
		IexAskSize            float64     `json:"iexAskSize"`
		MarketCap             float64     `json:"marketCap"`
		PeRatio               float64 `json:"peRatio"`
		Week52High            float64 `json:"week52High"`
		Week52Low             interface{} `json:"week52Low"`
		YtdChange             float64 `json:"ytdChange"`
}

// GetHandler handles to get a collection of listed companies

func GetCollectionHandler() {
	var sector string

	flag.StringVar(&sector, "sector", "Technology", "Technology sector")

	flag.Parse()

	url := fmt.Sprintf("https://api.iextrading.com/1.0/stock/market/collection/sector?collectionName=" + url.QueryEscape(sector) )

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

//	req.Header.Set("Authorization", "Bearer 8bf0df96-809b-41f7-8b04-4066fb889961")

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var record jsonData
	if jsonErr := json.Unmarshal(body, &record); jsonErr != nil {
		log.Println(jsonErr)
		return
	}

	fmt.Println("Jade said total record:", len(record))

	
	for a := 0; a < len(record); a++ {
			fmt.Println("no:",a, record[a].Symbol, record[a].CompanyName, "LatestPrice:" ,record[a].LatestPrice)
		}
}


func main(){
	GetCollectionHandler()
}

