package postgres

import (
	"chat/internal/types"
)

func (s *PostgresStorage) StoreMessage(msg types.Message, room string) {
	if _, ok := s.history[room]; !ok {
		s.history[room] = &types.MessageHistory{}
	}
	s.history[room] = &types.MessageHistory{
		Data:  append(s.history[room].Data, msg),
		Total: s.history[room].Total + 1,
	}
}

func (s *PostgresStorage) ListMessages(room string) *types.MessageHistory {
	return s.history[room]
}

func (s *PostgresStorage) EditMessage(msg types.Message, room string) {
	history := s.history[room]
	for i := range history.Data {
		if history.Data[i].ID == msg.ID {
			history.Data[i] = msg
		}
	}
}

func (s *PostgresStorage) GetMessage(uid uint, room string) types.Message {
	history := s.history[room]
	for i := range history.Data {
		if history.Data[i].ID == uid {
			return history.Data[i]
		}
	}
	return types.Message{}
}
