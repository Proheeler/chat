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
	EditMessage(msg types.Message, room string)
}

type ParticipantStorage interface {
	StoreParticipant(participant types.Client, room string)
	LoadParticipants(room string) types.ClientList
	DeleteParticipant(uid, room string)
	EditParticipant(participant types.Client, room string)
}

type RoomsStorage interface {
	ListRooms() []types.ShortRoomInfo
	GetRoom(room string) *types.Room
	CheckRoom(room string) bool
	AddRoom(room *types.Room)
	EditRoom(room *types.Room)
	LoadMessageHistory(room string) *types.MessageHistory
	DeleteRoom(room string)
}

type Searcher interface {
	Search(val, room string) []int
	GlobalSearch(val string) map[string][]int
}

type FileStorage interface {
	Upload()
	Download()
}
