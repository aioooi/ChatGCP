package main

import (
	"fmt"
	"net/http"

	"github.com/aioooi/goMicroservicesGCP/entity"
	"github.com/aioooi/goMicroservicesGCP/handlers"
	"github.com/gorilla/mux"
)

func main() {
	go entity.AwaitShouts()

	router := mux.NewRouter()

	router.Handle("/shouts", handlers.GetShoutHandler()).Methods("GET")
	router.Handle("/shouts", handlers.PostShoutHandler()).Methods("POST")

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server listening on port 8080.")
	server.ListenAndServe()
}
