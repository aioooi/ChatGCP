package main

import (
	"fmt"
	"net/http"

	"github.com/aioooi/goMicroservicesGCP/backend/entity"
	"github.com/aioooi/goMicroservicesGCP/backend/handlers"
	"github.com/gorilla/mux"
)

func main() {
	go entity.PersistShoutsFromQueue()

	router := mux.NewRouter()

	router.Handle("/shouts", handlers.GetShoutHandler()).Methods("GET")
	router.Handle("/shouts", handlers.PostShoutHandler()).Methods("POST")

	server := http.Server{
		Addr:    ":80",
		Handler: router,
	}

	fmt.Println("Server listening on port 80.")
	server.ListenAndServe()
}
