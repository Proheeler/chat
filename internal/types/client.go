package types

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ExternalID string
	Name       string
	Surname    string
	// Rooms pq.StringArray
	IsActive bool
	LastSeen time.Time
}

type ClientList struct {
	Total int
	Data  []Client
}

type Req struct {
	ID uint
}
