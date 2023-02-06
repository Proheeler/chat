package postgres

import (
	"chat/internal/types"
)

func (s *PostgresStorage) StoreMessage(msg types.Message, room string) {
	msg.Room = room
	s.db.Save(&msg)
}

func (s *PostgresStorage) ListMessages(room string, offset, limit int) *types.MessageHistory {
	msgs := []types.Message{}
	count := int64(0)
	s.db.Model(&types.Message{}).Where("room = ?", room).Count(&count)
	if count < int64(limit) {
		limit = int(count)
		offset = 0
	}
	if count < int64(offset*limit+limit) {
		offset = int(count) / limit
		limit = int(count) % limit
	}
	s.db.Order("id desc").Limit(limit).Offset(offset).Where("room = ?", room).Find(&msgs)
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
