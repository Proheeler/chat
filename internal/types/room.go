package types

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ShortRoomInfo
	Participants   pq.Int64Array `gorm:"type:integer[]"`
	PinnedMessages pq.Int64Array `gorm:"type:integer[]"`
}

type ShortRoomInfo struct {
	Name string
}

type ShortRoomInfoList struct {
	Total int
	Data  []ShortRoomInfo
}
