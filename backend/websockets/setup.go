package websockets

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/olahol/melody"
)

type MessageType string

const (
	// Build hierarchy of type and action/event/... if necessary:
	// PROTOCOL_ACTION MessageType = "action"
	// PROTOCOL_EVENT  MessageType = "event"
	// For now these suffice:
	ACTION_LOGIN        MessageType = "action-login"
	EVENT_USERS_UPDATED MessageType = "event-users-updated"
	PAYLOAD             MessageType = "payload"
)

type Message struct {
	Type    MessageType       `json:"type"`
	Payload map[string]string `json:"payload"`
}

type OnlineUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Maps "Sec-Webscocket-Key"s to user names and IDs of "logged in" users
var session = make(map[string]OnlineUser)

// First step of every melody handler:
// print some log line and return Sec-Websocket-Key
func initHandler(s *melody.Session, name string) string {
	fmt.Println("\n\n" + name + "\n----")
	k := s.Request.Header["Sec-Websocket-Key"][0]
	fmt.Println(k)
	return k
}

// Sets up the melody handlers
func Setup(mldy *melody.Melody) {
	mldy.HandleConnect(func(s *melody.Session) {
		initHandler(s, "handleConnect")
	})

	mldy.HandleDisconnect(func(s *melody.Session) {
		k := initHandler(s, "handleDisconnect")
		delete(session, k)
		fmt.Println(session)

		mldy.Broadcast([]byte("{\"type\":\"" + EVENT_USERS_UPDATED + "\"}"))
	})

	mldy.HandleMessage(func(s *melody.Session, msg []byte) {
		k := initHandler(s, "handleMessage")
		fmt.Println(string(msg))

		var m Message
		err := json.Unmarshal(msg, &m)
		if err != nil {
			fmt.Println(err)
			return
		}

		switch m.Type {
		case ACTION_LOGIN:
			// This is the only action so far, if more do another (outer, on
			// PROTOCOL_ACTION, PROTOCOL_EVENT, PAYLOAD) switch/case and handle type
			// of PROTOCOL_ACTION in inner switch/case, expecting action type such as
			// ACTION_LOGIN as part of the payload.
			session[k] = OnlineUser{string(m.Payload["user"]), uuid.New().String()}
			fmt.Println(session)

			mldy.Broadcast([]byte("{\"type\":\"" + EVENT_USERS_UPDATED + "\"}"))

		case PAYLOAD:
			// TODO maybe broadcast _to others_ to handle "own" messages without user
			// IDs?
			mldy.Broadcast(msg)
		}
	})
}
