package main

import (
	"api-postgresql/configs"
	"api-postgresql/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)
	r.Delete("/{id}", handlers.Delete)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
