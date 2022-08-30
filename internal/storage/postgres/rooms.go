package postgres

import (
	"chat/internal/types"
)

func (s *PostgresStorage) ListRooms() []types.ShortRoomInfo {
	keys := make([]types.ShortRoomInfo, len(s.rooms))
	i := 0
	for k := range s.rooms {
		keys[i] = s.rooms[k].ShortRoomInfo
		i++
	}
	return keys
}

func (s *PostgresStorage) CheckRoom(room string) bool {
	_, ok := s.rooms[room]
	return ok
}

func (s *PostgresStorage) GetRoom(room string) *types.Room {
	rm, ok := s.rooms[room]
	if !ok {
		return nil
	}
	return &rm
}

func (s *PostgresStorage) AddRoom(room *types.Room) {
	s.rooms[room.Name] = *room
	s.history[room.Name] = &types.MessageHistory{
		Total: 0,
		Data:  []types.Message{},
	}
}

func (s *PostgresStorage) EditRoom(prevName string, room *types.Room) {
	s.rooms[room.Name] = types.Room{
		ShortRoomInfo: types.ShortRoomInfo{
			Name: room.Name,
		},
		Participants:   s.rooms[prevName].Participants,
		PinnedMessages: s.rooms[prevName].PinnedMessages,
	}
	delete(s.rooms, prevName)
}

func (s *PostgresStorage) DeleteRoom(room string) {
	delete(s.rooms, room)
}
