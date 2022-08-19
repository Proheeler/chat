package simple

import (
	"chat/internal/types"
	"time"
)

func (s *SimpleStorage) StoreParticipant(patricipant types.Client, room string) {
	rm := s.rooms[room]
	rm.Participants = append(s.rooms[room].Participants, patricipant)
	s.rooms[room] = rm
}
func (s *SimpleStorage) ListParticipants(room string) types.ClientList {
	return types.ClientList{
		Data:  s.rooms[room].Participants,
		Total: len(s.rooms[room].Participants),
	}
}

func (s *SimpleStorage) GetParticipant(id, room string) types.Client {
	var client types.Client
	for i := range s.rooms[room].Participants {
		if s.rooms[room].Participants[i].ID == id || s.rooms[room].Participants[i].ExternalID == id {
			client = s.rooms[room].Participants[i]
			break
		}
	}
	return client
}

func (s *SimpleStorage) EditParticipant(participant types.Client, room string) {
	pl := s.rooms[room].Participants
	for i := range pl {
		if pl[i].ID == participant.ID {
			s.rooms[room].Participants[i] = participant
			s.rooms[room].Participants[i].UpdatedAt = time.Now()
		}
	}
}

func (s *SimpleStorage) DeleteParticipant(uid, room string) {
	pl := s.rooms[room].Participants
	rm := s.rooms[room]
	for i := range pl {
		if pl[i].ID == uid {
			rm.Participants = RemoveIndex(pl, i)
			break
		}
	}
	s.rooms[room] = rm
}

func RemoveIndex(s []types.Client, index int) []types.Client {
	return append(s[:index], s[index+1:]...)
}
