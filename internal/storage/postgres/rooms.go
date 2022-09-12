package postgres

import (
	"chat/internal/types"
)

func (s *PostgresStorage) ListRooms() *types.ShortRoomInfoList {
	rooms := []types.Room{}
	s.db.Find(&rooms)
	return &types.ShortRoomInfoList{
		Total: len(rooms),
		Data:  rooms,
	}
}

func (s *PostgresStorage) CheckRoom(room string) bool {
	rm := &types.Room{}
	tx := s.db.Where("name = ?", room).Find(rm)

	return tx.Error == nil
}

func (s *PostgresStorage) GetRoom(room string) *types.Room {
	rm := &types.Room{}
	tx := s.db.Where("name = ?", room).Find(rm)
	if tx.Error != nil {
		return nil
	}
	return rm
}

func (s *PostgresStorage) AddRoom(room *types.Room) {
	s.db.Save(room)
}

func (s *PostgresStorage) EditRoom(prevName string, room *types.Room) {
	s.db.Save(room)
}

func (s *PostgresStorage) DeleteRoom(room string) {
	s.db.Where("name = ?", room).Delete(&types.Room{})
}
