package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//type ticjer info
type TickerInfo struct {
	High      string `json:"high"`
	Last      string `json:"last"`
	Timestamp string `json:"timestamp"`
	Bid       string `json:"bid"`
	Vwap      string `json:"vwap"`
	Volume    string `json:"volume"`
	Low       string `json:"low"`
	Ask       string `json:"ask"`
	Open      string `json:"open"`
}
type TickerManagerService struct {
	httpClient *http.Client
}

//Service represents the bitsmap app to get ticket
func NewTickerManagerService() *TickerManagerService {
	return &TickerManagerService{
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// This gets ticket info from bitsmap
func (s *TickerManagerService) GetTicker() TickerInfo {
	bitsmapApiGetTickerUrl := "https://www.bitstamp.net/api/v2/ticker/btcusd/"
	fmt.Println(bitsmapApiGetTickerUrl)
	resp, err := s.httpClient.Get(bitsmapApiGetTickerUrl)

	if err != nil {
		return TickerInfo{}
	}
	if resp.StatusCode != 200 {
		return TickerInfo{}
	}

	defer resp.Body.Close()
	body := new(TickerInfo)
	err = json.NewDecoder(resp.Body).Decode(body)
	if err != nil {
		fmt.Printf("Error when decoding body %v", err)
		return TickerInfo{}
	}

	ticker := *body
	fmt.Println(ticker)
	return ticker
}
