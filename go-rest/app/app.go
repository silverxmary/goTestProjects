package app

import (
	"net/http"
	"time"

	"git.manomano.tech/mariellys.soto/go-rest/app/database"
	"github.com/gorilla/mux"
)

type App struct {
	Router     *mux.Router     //start router
	DB         database.PostDB //instance to inferface
	httpClient *http.Client
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/posts", a.CreatePostHandler()).Methods("POST")
	a.Router.HandleFunc("/api/posts", a.GetPostsHandler()).Methods("GET")
	a.Router.HandleFunc("/api/ticker", a.GetTickerHandler()).Methods("GET")
}
