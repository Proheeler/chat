package storage

import (
	"chat/internal/types"
)

type Storage interface {
	MessageStorage
	ParticipantStorage
	RoomsStorage
	Searcher
}

type MessageStorage interface {
	StoreMessage(msg types.Message, room string)
	LoadMessageHistory(room string) *types.MessageHistory
	EditMessage(msg types.Message, room string)
}

type ParticipantStorage interface {
	StoreParticipant(participant types.Person, room string)
	LoadParticipants(room string) types.PersonList
	DeleteParticipant(uid, room string)
}

type RoomsStorage interface {
	ListRooms() []string
	CheckRoom(room string) bool
	AddRoom(room string)
}

type Searcher interface {
	Search(val, room string) []int
	GlobalSearch(val string) map[string][]int
}

type FileStorage interface {
	Upload()
	Download()
}
