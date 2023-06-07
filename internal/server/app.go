package server

import (
	"github.com/DmitryOdintsov/workingWithGit/internal/server/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

var addr = ":8282"

type App struct {
	router chi.Router
}

func Run(h *handlers.Handler) {
	log.Printf("starting server on %s", addr)

	router := chi.NewRouter()
	router.Get("/api", h.HandleConnection)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalln(err)
	}

}
