package types

import "time"

type Client struct {
	ID         string
	ExternalID string
	Name       string
	Surname    string
	CreateAt   time.Time `json:"create_at,omitempty"`
	UpdateAt   time.Time `json:"update_at,omitempty"`
}

type ClientList struct {
	Total int
	Data  []Client
}
