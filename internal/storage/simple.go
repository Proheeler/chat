package storage

import (
	"chat/internal/types"
	"strings"

	"github.com/google/uuid"
)

type SimpleStorage struct {
	rooms map[string]types.Room
}

func NewSimpleStorage() *SimpleStorage {
	return &SimpleStorage{
		rooms: map[string]types.Room{},
	}
}

func (s *SimpleStorage) StoreMessage(msg types.Message, room string) {
	if _, ok := s.rooms[room]; !ok {
		s.rooms[room] = types.Room{
			History: &types.MessageHistory{},
		}
	}
	s.rooms[room] = types.Room{
		ID:             s.rooms[room].ID,
		Name:           s.rooms[room].Name,
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

func (s *SimpleStorage) Search(val, room string) []int {
	ret := []int{}
	for i := range s.rooms[room].History.Data {
		if strings.Contains(s.rooms[room].History.Data[i].Data, val) {
			ret = append(ret, i)
		}
	}
	return ret
}
func (s *SimpleStorage) GlobalSearch(val string) map[string][]int {
	ret := map[string][]int{}
	for k := range s.rooms {
		for i := range s.rooms[k].History.Data {
			if strings.Contains(s.rooms[k].History.Data[i].Data, val) {
				ret[k] = append(ret[k], i)
			}
		}
	}

	return ret
}

func (s *SimpleStorage) ListRooms() []string {
	keys := make([]string, len(s.rooms))
	i := 0
	for k := range s.rooms {
		keys[i] = k
		i++
	}
	return keys
}

func (s *SimpleStorage) CheckRoom(room string) bool {
	_, ok := s.rooms[room]
	return ok
}

func (s *SimpleStorage) AddRoom(room string) {
	s.rooms[room] = types.Room{
		ID:             uuid.New().String(),
		Name:           room,
		Participants:   []types.Client{},
		PinnedMessages: []string{},
		History:        &types.MessageHistory{},
	}
}

func (s *SimpleStorage) StoreParticipant(patricipant types.Client, room string) {
	rm := s.rooms[room]
	rm.Participants = append(s.rooms[room].Participants, patricipant)
	s.rooms[room] = rm
}
func (s *SimpleStorage) LoadParticipants(room string) types.ClientList {
	return types.ClientList{
		Data:  s.rooms[room].Participants,
		Total: len(s.rooms[room].Participants),
	}
}

func (s *SimpleStorage) EditParticipant(participant types.Client, room string) {
	// pl := s.rooms[room].Participants
	// for i := range pl {
	// 	if pl[i].ID == participant.ID {
	// 		s.rooms[room].Participants[i] = participant
	// 	}
	// }
}

func (s *SimpleStorage) DeleteParticipant(uid, room string) {
	pl := s.rooms[room].Participants
	rm := s.rooms[room]
	for i := range pl {
		if pl[i].ID == uid {
			rm.Participants = RemoveIndex(pl, i)
			break
		}
	}
	s.rooms[room] = rm
}

func RemoveIndex(s []types.Client, index int) []types.Client {
	return append(s[:index], s[index+1:]...)
}
