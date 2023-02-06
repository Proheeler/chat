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
	GetMessage(uid uint, room string) types.Message
	ListMessages(room string) *types.MessageHistory
}

type ParticipantStorage interface {
	StoreParticipant(participant types.Client)
	DeleteParticipant(uid uint)
	GetParticipant(uid uint) types.Client
	EditParticipant(participant types.Client)
	ListParticipants() types.ClientList
}

type RoomsStorage interface {
	ListRooms() *types.ShortRoomInfoList
	GetRoom(room string) *types.Room
	CheckRoom(room string) bool
	AddRoom(room *types.Room) error
	EditRoom(prevName string, room *types.Room) error
	DeleteRoom(room string)
	AddParticipantInRoom(patricipant uint, room string)
	ListParticipantsInRoom(room string) types.ClientList
	DeleteParticipantInRoom(uid uint, room string)
}

type Searcher interface {
	Search(val, room string) []types.Message
	GlobalSearch(val string) map[string][]types.Message
}

type FileStorage interface {
	Upload()
	Download()
}
