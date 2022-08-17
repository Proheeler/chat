package simple

import "chat/internal/types"

func (s *SimpleStorage) StoreMessage(msg types.Message, room string) {
	if _, ok := s.rooms[room]; !ok {
		s.rooms[room] = types.Room{
			History: &types.MessageHistory{},
		}
	}
	s.rooms[room] = types.Room{
		ShortRoomInfo: types.ShortRoomInfo{
			ID:   s.rooms[room].ID,
			Name: s.rooms[room].Name,
		},
		Participants:   s.rooms[room].Participants,
		PinnedMessages: s.rooms[room].PinnedMessages,
		History: &types.MessageHistory{
			Data:  append(s.rooms[room].History.Data, msg),
			Total: s.rooms[room].History.Total + 1,
		},
	}
}

func (s *SimpleStorage) LoadMessageHistory(room string) *types.MessageHistory {
	return s.rooms[room].History
}

func (s *SimpleStorage) EditMessage(msg types.Message, room string) {
	history := s.rooms[room].History
	for i := range history.Data {
		if history.Data[i].ID == msg.ID {
			history.Data[i] = msg
		}
	}
}
