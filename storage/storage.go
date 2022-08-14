package storage

import (
	"chat/types"
)

type Storage interface {
	MessageStorage
	ParticipantStorage
}

type MessageStorage interface {
	StoreMessage(msg types.Message, room string)
	LoadMessageHistory(room string) types.MessageHistory
}

type ParticipantStorage interface {
	StoreParticipant(patricipant types.Person, room string)
	LoadParticipants(room string) types.PersonList
	DeleteParticipant(uid, room string)
}
