package types

import "time"

type Message struct {
	ID           string
	Sender       string
	ReplyTo      string
	Data         string
	CreationTime time.Time
	EditTime     time.Time
	Attachments  []Media
}

type MessageHistory struct {
	Total int
	Data  []Message
}
