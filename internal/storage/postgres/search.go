package postgres

import (
	"chat/internal/types"
)

func (s *PostgresStorage) Search(val, room string) []types.Message {
	cls := []types.Message{}
	s.db.Where("data LIKE ? AND room = ?", "%"+val+"%", room).Find(&cls)
	return cls
}
func (s *PostgresStorage) GlobalSearch(val string) map[string][]types.Message {
	ret := map[string][]types.Message{}
	cls := []types.Message{}
	s.db.Where("data LIKE ?", "%"+val+"%").Find(&cls)

	for k := range cls {
		ret[cls[k].Room] = append(ret[cls[k].Room], cls[k])
	}
	return ret
}
