package types

import "time"

type Client struct {
	ID         string
	ExternalID string
	Name       string
	Surname    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Rooms      []string
	IsActive   bool
	LastSeen   time.Time
}

type ClientList struct {
	Total int
	Data  []Client
}

type Req struct {
	ID string
}
