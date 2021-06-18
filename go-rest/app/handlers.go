package app

import (
	"fmt"
	"log"
	"net/http"

	"git.manomano.tech/mariellys.soto/go-rest/app/models"
	"git.manomano.tech/mariellys.soto/go-rest/helpers"

)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Post API")
	}
}

func (a *App) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := helpers.Parse(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			helpers.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		// Create the post
		/*p := &models.Post{
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		}*/
		p := helpers.MapToPost(&req)

		// Save in DB
		err = a.DB.CreatePost(&p)
		if err != nil {
			log.Printf("Cannot save post in DB. err=%v \n", err)
			helpers.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := helpers.MapPostToJSON(&p)
		helpers.SendResponse(w, r, resp, http.StatusOK)

	}
}

func (a *App) GetPostsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := a.DB.GetPosts()
		if err != nil {
			log.Printf("Cannot get posts, err=%v \n", err)
			helpers.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonPost, len(posts))
		for idx, post := range posts {
			resp[idx] = helpers.MapPostToJSON(post)
		}

		helpers.SendResponse(w, r, resp, http.StatusOK)
	}
}
