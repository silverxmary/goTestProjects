package api

import (
	"fmt"
	"net/http"
)

// Ticker Handler interface ...
type TickerManagerService interface {
	GetTicker() TickerInfoDto
}

//ticker service
type TickerManageHandler struct {
	svc TickerManagerService
}

//Type ticker info
type TickerInfoDto struct {
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

//Func that gets ticket info from another service and then return info
func (t *TickerManageHandler) GetTickerInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the local bitsmap!")
	fmt.Println("Endpoint Hit: GetTickerInfo")
	fmt.Println(t.svc.GetTicker())
	//return t.svc.GetTicker()
	//return NewTickerManagerService.GetTicker()
}
