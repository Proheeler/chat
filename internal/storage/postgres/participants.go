package postgres

import (
	"chat/internal/types"
)

func (s *PostgresStorage) StoreParticipant(patricipant types.Client) {
	s.db.Save(&patricipant)
}

func (s *PostgresStorage) GetParticipant(id uint) types.Client {
	cl := &types.Client{}
	s.db.Where("external_id = ?", id).Find(cl)
	return *cl
}

func (s *PostgresStorage) ListParticipants() types.ClientList {
	cls := []types.Client{}
	s.db.Find(&cls)
	return types.ClientList{
		Total: len(cls),
		Data:  cls,
	}
}

func (s *PostgresStorage) EditParticipant(participant types.Client) {
	s.db.Save(&participant)
}

func (s *PostgresStorage) DeleteParticipant(uid uint) {
	s.db.Delete(&types.Client{}, uid)
}
