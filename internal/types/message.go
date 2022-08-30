package types

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Sender      string
	ReplyTo     string
	Data        string
	Attachments pq.Int64Array `gorm:"type:integer[]"`
}

type MessageHistory struct {
	Total int
	Data  []Message
}
