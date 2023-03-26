package entity

// A simple shout post
type Shout struct {
	Id        string `json:"id"`
	Author    string `json:"author"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}
