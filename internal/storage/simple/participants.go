package simple

import (
	"chat/internal/types"
)

func (s *SimpleStorage) StoreParticipant(patricipant types.Client) {
	s.clients[int64(patricipant.ID)] = patricipant
}

func (s *SimpleStorage) GetParticipant(id int64) types.Client {
	return s.clients[id]
}

func (s *SimpleStorage) ListParticipants() types.ClientList {
	var parts []types.Client
	for i := range s.clients {
		parts = append(parts, s.clients[i])
	}
	return types.ClientList{
		Data:  parts,
		Total: len(parts),
	}
}

func (s *SimpleStorage) EditParticipant(participant types.Client) {
	pl := s.clients
	for i := range pl {
		if pl[i].ID == participant.ID {
			s.clients[int64(participant.ID)] = participant
			break
		}
	}
}

func (s *SimpleStorage) DeleteParticipant(uid int64) {
	delete(s.clients, uid)
}
