package main

import (
	"fmt"
	"net/http"

	"github.com/aioooi/goMicroservicesGCP/backend/entity"
	"github.com/aioooi/goMicroservicesGCP/backend/handlers"
	"github.com/aioooi/goMicroservicesGCP/backend/websockets"
	"github.com/gorilla/mux"
	"github.com/olahol/melody"
)

func main() {
	go entity.PersistShoutsFromQueue()

	router := mux.NewRouter()

	// Static files:
	router.Handle("/", http.FileServer(http.Dir("./public")))

	// API
	router.Handle("/shouts", handlers.GetShoutHandler()).Methods("GET")
	router.Handle("/shouts", handlers.PostShoutHandler()).Methods("POST")

	// Websockets
	mel := melody.New()

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		mel.HandleRequest(w, r)
	})

	websockets.Setup(mel)

	server := http.Server{
		Addr:    ":80",
		Handler: router,
	}

	fmt.Println("Server listening on port 80.")
	server.ListenAndServe()
}
