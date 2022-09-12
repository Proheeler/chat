package postgres

import (
	"chat/internal/types"
)

func (s *PostgresStorage) AddParticipantInRoom(patricipant uint, room string) {
	rm := &types.Room{}
	tx := s.db.Where("name = ?", room).Find(rm)
	if tx.Error != nil {
		return
	}
	isFound := false
	for i := range rm.Participants {
		if rm.Participants[i] == int64(patricipant) {
			isFound = true
			break
		}
	}
	if !isFound {
		rm.Participants = append(rm.Participants, int64(patricipant))
		s.db.Save(rm)
	}
}
func (s *PostgresStorage) ListParticipantsInRoom(room string) types.ClientList {
	var parts []types.Client
	rm := &types.Room{}
	tx := s.db.Where("name = ?", room).Find(rm)
	if tx.Error != nil {
		return types.ClientList{
			Data:  parts,
			Total: len(parts),
		}
	}
	cls := []uint{}
	for i := range rm.Participants {
		cls = append(cls, uint(rm.Participants[i]))
	}
	s.db.Where("id IN ?", cls).Find(&parts)
	return types.ClientList{
		Data:  parts,
		Total: len(parts),
	}
}

func (s *PostgresStorage) DeleteParticipantInRoom(uid uint, room string) {
	rm := &types.Room{}
	tx := s.db.Where("name = ?", room).Find(rm)
	if tx.Error != nil {
		return
	}
	for i := range rm.Participants {
		if rm.Participants[i] == int64(uid) {
			rm.Participants = RemoveIndex(rm.Participants, i)
			s.db.Save(rm)
		}
	}
}

func RemoveIndex(s []int64, index int) []int64 {
	return append(s[:index], s[index+1:]...)
}
