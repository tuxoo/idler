package entity

import "time"

type Message struct {
	Sender string
	SentAt time.Time
	Text   string
}
