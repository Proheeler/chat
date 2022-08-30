package simple

import (
	"chat/internal/types"
)

type SimpleStorage struct {
	rooms   map[string]types.Room
	history map[string]*types.MessageHistory
	clients map[int64]types.Client
}

func NewSimpleStorage() *SimpleStorage {
	return &SimpleStorage{
		rooms:   map[string]types.Room{},
		history: map[string]*types.MessageHistory{},
		clients: map[int64]types.Client{},
	}
}
