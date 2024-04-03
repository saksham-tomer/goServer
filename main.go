package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {
	port := godotenv.Load()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Ues(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.write([]byte("root."))
	})

	http.listenAndServe(port, r)
}
