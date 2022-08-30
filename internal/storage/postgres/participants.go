package postgres

import (
	"chat/internal/types"
)

func (s *PostgresStorage) StoreParticipant(patricipant types.Client) {
	s.clients[patricipant.ID] = patricipant
}

func (s *PostgresStorage) GetParticipant(id uint) types.Client {
	return s.clients[id]
}

func (s *PostgresStorage) ListParticipants() types.ClientList {
	var parts []types.Client
	for i := range s.clients {
		parts = append(parts, s.clients[i])
	}
	return types.ClientList{
		Data:  parts,
		Total: len(parts),
	}
}

func (s *PostgresStorage) EditParticipant(participant types.Client) {
	pl := s.clients
	for i := range pl {
		if pl[i].ID == participant.ID {
			s.clients[participant.ID] = participant
			break
		}
	}
}

func (s *PostgresStorage) DeleteParticipant(uid uint) {
	delete(s.clients, uid)
}
