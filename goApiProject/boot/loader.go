package boot

import (
	"fmt"

	"git.manomano.tech/mariellys.soto/goApiProject/http"
	//"git.manomano.tech/mariellys.soto/gotestproject/tickers"
)

// This function starts the application
func AppRun() {

	fmt.Println("App starting")
	http.HandleRequests()
	//fmt.Println("Get Ticker")
	//tickers.NewTickerManagerService().GetTicker()
}
