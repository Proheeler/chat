package types

import "time"

type Message struct {
	Sender      string
	ReplyTo     string
	Data        string
	Time        time.Time
	Attachments []Media
}

type MessageHistory struct {
	Total int
	Data  []Message
}
