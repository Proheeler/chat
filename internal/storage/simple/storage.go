package simple

import (
	"chat/internal/types"
)

type SimpleStorage struct {
	rooms map[string]types.Room
}

func NewSimpleStorage() *SimpleStorage {
	return &SimpleStorage{
		rooms: map[string]types.Room{},
	}
}
