package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	"git.manomano.tech/mariellys.soto/go-rest/app/models"
)

func Parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func SendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}

func MapPostToJSON(p *models.Post) models.JsonPost {
	return models.JsonPost{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
		Author:  p.Author,
	}
}

func MapToPost(p *models.PostRequest) models.Post {
	return models.Post{
		Title:   p.Title,
		Content: p.Content,
		Author:  p.Author,
	}
}

func MapTickerInfoToTickerResponse(t *models.TickerInfo) models.TickerInfoResponse {
	return models.TickerInfoResponse{
		High:      t.High,
		Last:      t.Last,
		Timestamp: t.Timestamp,
	}
}
