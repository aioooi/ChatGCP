package entity

import (
	"time"

	"github.com/google/uuid"
)

// A simple shout post
type Shout struct {
	Id        string `json:"id"`
	Author    string `json:"author"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// The "database":
var shouts []*Shout

func AllShouts() []*Shout {
	return shouts
}

func PostShout(author, message string) {
	shouts = append(shouts, &Shout{
		Id:        uuid.New().String(),
		Author:    author,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC822),
	})
}
