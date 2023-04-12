package websockets

import (
	"fmt"
	"net/http"

	"github.com/olahol/melody"
)

// Prints relevant parts of an http.Request
func printRequest(r *http.Request) {
	// for k, v := range h.Header {
	// 	fmt.Println(k)
	// 	fmt.Println(v)
	// }
	fmt.Println(r.Header["Sec-Websocket-Key"])
}

// Sets up the melody handlers
func Setup(m *melody.Melody) {
	m.HandleConnect(func(s *melody.Session) {
		fmt.Println("\n\nHandleConnect\n------")
		printRequest(s.Request)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("\n\nHandleDisconnect\n------")
		printRequest(s.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println("\n\nHandleMessage\n------")
		m.Broadcast(msg)
		printRequest(s.Request)
	})
}

