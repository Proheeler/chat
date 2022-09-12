package simple

import (
	"chat/internal/types"
)

func (s *SimpleStorage) ListRooms() *types.ShortRoomInfoList {
	keys := make([]types.Room, len(s.rooms))
	i := 0
	for k := range s.rooms {
		keys[i] = s.rooms[k]
		i++
	}
	return &types.ShortRoomInfoList{
		Total: len(s.rooms),
		Data:  keys,
	}
}

func (s *SimpleStorage) CheckRoom(room string) bool {
	_, ok := s.rooms[room]
	return ok
}

func (s *SimpleStorage) GetRoom(room string) *types.Room {
	rm, ok := s.rooms[room]
	if !ok {
		return nil
	}
	return &rm
}

func (s *SimpleStorage) AddRoom(room *types.Room) {
	s.rooms[room.Name] = *room
	s.history[room.Name] = &types.MessageHistory{
		Total: 0,
		Data:  []types.Message{},
	}
}

func (s *SimpleStorage) EditRoom(prevName string, room *types.Room) {
	s.rooms[room.Name] = types.Room{
		ShortRoomInfo: types.ShortRoomInfo{
			Name: room.Name,
		},
		Participants:   s.rooms[prevName].Participants,
		PinnedMessages: s.rooms[prevName].PinnedMessages,
	}
	delete(s.rooms, prevName)
}

func (s *SimpleStorage) DeleteRoom(room string) {
	delete(s.rooms, room)
}
