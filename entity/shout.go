package entity

import (
	"math"
	"time"

	"github.com/google/uuid"
)

// A simple shout post:
type Shout struct {
	Id        string `json:"id"`
	Author    string `json:"author"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// The "database":
var shouts []*Shout
var shoutQueue = make(chan *Shout)

const MAX_SHOUTS int = 20

func AwaitShouts() {
	for {
		shouts = shouts[int(math.Max(float64(len(shouts)-MAX_SHOUTS), 0.0)):]
		shouts = append(shouts, <-shoutQueue)
	}
}

func AllShouts() *[]*Shout {
	return &shouts
}

func PostShout(author, message string) {
	shoutQueue <- &Shout{
		Id:        uuid.New().String(),
		Author:    author,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
