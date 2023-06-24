package main

import (
	"API/config"
	"API/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Route("/todo", handlers.TodoRouter)

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetAPIConfig().Port), router)
}
