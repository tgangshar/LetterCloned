package handlers

import (
	"goth/internal/middleware"
	"goth/internal/store"
	"goth/internal/templates"
	"net/http"

	"fmt"
	"log"
	"github.com/ryanbradynd05/go-tmdb"
)

var tmdbAPI *tmdb.TMDb

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value(middleware.UserKey).(*store.User)

	config := tmdb.Config{
		APIKey: "c80a280df643024ba7acafa7f74f269b",
		Proxies: nil,
		UseProxy: false,
	}
	tmdbAPI = tmdb.Init(config)

	fightClubInfo, tmdbErr := tmdbAPI.GetMovieInfo(550, nil)
	fightClubJson, tmdbErr := tmdb.ToJSON(fightClubInfo)

	if tmdbErr != nil {
		log.Fatal(tmdbErr)
	}

	fmt.Println(fightClubInfo)

	if !ok {
		c := templates.GuestIndex()
		err := templates.Layout(c, fightClubJson).Render(r.Context(), w)

		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}

		return
	}

	c := templates.Index(user.Email)
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
