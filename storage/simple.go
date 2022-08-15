package storage

import (
	"chat/types"
	"strings"
)

type SimpleStorage struct {
	History      map[string]types.MessageHistory
	Participants map[string]types.PersonList
}

func NewSimpleStorage() *SimpleStorage {
	return &SimpleStorage{
		History: map[string]types.MessageHistory{},
	}
}

func (s *SimpleStorage) StoreMessage(msg types.Message, room string) {
	s.History[room] = types.MessageHistory{
		Data:  append(s.History[room].Data, msg),
		Total: s.History[room].Total + 1,
	}
}

func (s *SimpleStorage) LoadMessageHistory(room string) types.MessageHistory {
	return s.History[room]
}

func (s *SimpleStorage) Search(val, room string) []int {
	ret := []int{}
	for i := range s.History[room].Data {
		if strings.Contains(s.History[room].Data[i].Data, val) {
			ret = append(ret, i)
		}
	}
	return ret
}
func (s *SimpleStorage) GlobalSearch(val string) map[string][]int {
	ret := map[string][]int{}
	for k := range s.History {
		for i := range s.History[k].Data {
			if strings.Contains(s.History[k].Data[i].Data, val) {
				ret[k] = append(ret[k], i)
			}
		}
	}

	return ret
}

func (s *SimpleStorage) ListRooms() []string {
	keys := make([]string, len(s.History))
	i := 0
	for k := range s.History {
		keys[i] = k
		i++
	}
	return keys
}

func (s *SimpleStorage) StoreParticipant(patricipant types.Person, room string) {
	s.Participants[room] = types.PersonList{
		Data:  append(s.Participants[room].Data, patricipant),
		Total: s.History[room].Total + 1,
	}
}
func (s *SimpleStorage) LoadParticipants(room string) types.PersonList {
	return s.Participants[room]
}
func (s *SimpleStorage) DeleteParticipant(uid, room string) {
	pl := s.Participants[room]
	for i := range pl.Data {
		if pl.Data[i].ClientID == uid {
			RemoveIndex(pl.Data, i)
			break
		}
	}
}

func RemoveIndex(s []types.Person, index int) []types.Person {
	return append(s[:index], s[index+1:]...)
}
