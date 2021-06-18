package app

import (
	"encoding/json"
	"log"
	"net/http"

	"git.manomano.tech/mariellys.soto/go-rest/app/models"
	"git.manomano.tech/mariellys.soto/go-rest/helpers"
)

/*type TickerHandler struct {
	app        App
	httpClient *http.Client
}

//Service represents the bitsmap app to get ticket
func NewTickerHandler() *TickerHandler {
	return &TickerHandler{
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}*/
func (a *App) GetTickerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bitsmapApiGetTickerUrl := "https://www.bitstamp.net/api/v2/ticker/btcusd/"
		//fmt.Println(bitsmapApiGetTickerUrl)
		resp, err := a.httpClient.Get(bitsmapApiGetTickerUrl)

		if err != nil {
			log.Printf("Cannot get tickers, err=%v \n", err)
			helpers.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
		if resp.StatusCode != 200 {
			log.Printf("Cannot get tickers  err=%v \n", err)
			helpers.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		defer resp.Body.Close()
		body := new(models.TickerInfo)
		err = json.NewDecoder(resp.Body).Decode(body)
		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			helpers.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		ticker := body
		//fmt.Println(ticker)
		tickerResponse := helpers.MapTickerInfoToTickerResponse(ticker)

		helpers.SendResponse(w, r, tickerResponse, http.StatusOK)

	}

}
