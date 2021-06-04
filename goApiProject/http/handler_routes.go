package http

import (
	"log"
	"net/http"

	"git.manomano.tech/mariellys.soto/goApiProject/api"
)

//Func manages the urls and apis registration
func HandleRequests() {
	http.HandleFunc("/ticker", api.GetTickerInfo())
	log.Fatal(http.ListenAndServe(":10000", nil))
}
