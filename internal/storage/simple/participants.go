package simple

import (
	"chat/internal/types"
)

func (s *SimpleStorage) StoreParticipant(participant types.Client) {
	s.clients[participant.ID] = participant
}

func (s *SimpleStorage) GetParticipant(uid uint) types.Client {
	return s.clients[uid]
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
			s.clients[participant.ID] = participant
			break
		}
	}
}

func (s *SimpleStorage) DeleteParticipant(uid uint) {
	delete(s.clients, uid)
}
