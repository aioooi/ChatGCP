package websockets

import (
	"encoding/json"
	"fmt"

	"github.com/olahol/melody"
)

type MessageType string

const (
	// PROTOCOL_ACTION MessageType = "action"
	ACTION_LOGIN MessageType = "action-login"
	// PROTOCOL_EVENT  MessageType = "event"
	EVENT_LOGIN  MessageType = "event-login"
	EVENT_LOGOUT MessageType = "event-logout"
	PAYLOAD      MessageType = "payload"
)

type Message struct {
	Type    MessageType `json:"type"`
	Payload []byte      `json:"payload"`
}

// Maps "Sec-Webscocket-Key"s to user names of "logged in" users (or "" for
// connected sessions with undefined user name (== not "logged in")).
var session = make(map[string]string)

// First step of every melody handler:
// print some log line and return Sec-Websocket-Key
func initHandler(s *melody.Session, name string) string {
	fmt.Println("\n\n" + name + "\n----")
	k := s.Request.Header["Sec-Websocket-Key"][0]
	fmt.Println(k)
	return k
}

// Sets up the melody handlers
func Setup(m *melody.Melody) {
	m.HandleConnect(func(s *melody.Session) {
		k := initHandler(s, "handleConnect")
		session[k] = ""
		fmt.Println(session)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		k := initHandler(s, "handleDisconnect")
		delete(session, k)
		fmt.Println(session)

		m.Broadcast([]byte("{\"type\":" + EVENT_LOGOUT + "}"))
	})

	m.HandleMessage(func(s *melody.Session, _msg []byte) {
		k := initHandler(s, "handleMessage")

		var msg Message
		err := json.Unmarshal(_msg, &msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		switch msg.Type {
		case ACTION_LOGIN:
			// This is the only action so far, if more do another (outer, on
			// PROTOCOL_ACTION, PROTOCOL_EVENT, PAYLOAD) switch/case and handle type
			// of PROTOCOL_ACTION in inner switch/case, expecting action type such as
			// ACTION_LOGIN as part of the payload.
			session[k] = string(msg.Payload)
			fmt.Println(session)

			m.Broadcast([]byte("{\"type\":" + EVENT_LOGIN + "}"))

		case PAYLOAD:
			m.Broadcast(_msg)
		}
	})
}
