package main

import "time"

type Message struct {
	From string
	To   string
	Data string
	Time time.Time
}

type MessageHistory struct {
	Total int
	Data  []Message
}
