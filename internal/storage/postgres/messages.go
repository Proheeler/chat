package postgres

import (
	"chat/internal/types"
)

func (s *PostgresStorage) StoreMessage(msg types.Message, room string) {
	msg.Room = room
	s.db.Save(&msg)
}

func (s *PostgresStorage) ListMessages(room string) *types.MessageHistory {
	msgs := []types.Message{}
	s.db.Where("room = ?", room).Find(&msgs)
	return &types.MessageHistory{
		Total: len(msgs),
		Data:  msgs,
	}
}

func (s *PostgresStorage) EditMessage(msg types.Message, room string) {
	s.db.Save(&msg)
}
func (s *PostgresStorage) GetMessage(uid uint, room string) types.Message {
	msg := &types.Message{}
	s.db.First(msg, uid)
	return *msg
}
