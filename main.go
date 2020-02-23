package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	_ "github.com/GoogleCloudPlatform/berglas/pkg/auto"
)

type Message struct {
	Status int    `json:"status"`
	Text   string `json:"text"`
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		m := &Message{
			Status: http.StatusOK,
			Text:   os.Getenv("MESSAGE"),
		}
		render.JSON(w, r, m)
	})

	r.Get("/debug", func(w http.ResponseWriter, r *http.Request) {
		m := &Message{
			Status: http.StatusOK,
			Text:   os.Getenv("SECRET"),
		}
		render.JSON(w, r, m)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}
