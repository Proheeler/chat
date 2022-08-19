package types

import "time"

type Room struct {
	ShortRoomInfo
	Participants   []Client
	PinnedMessages []string
}

type ShortRoomInfo struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ShortRoomInfoList struct {
	Total int
	Data  []ShortRoomInfo
}
