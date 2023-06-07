package model

import "time"

type Message struct {
	Author    string    `json:"author,omitempty"`
	Text      string    `json:"text,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
