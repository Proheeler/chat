package simple

import (
	"chat/internal/types"
	"strings"
)

func (s *SimpleStorage) Search(val, room string) []types.Message {
	ret := []types.Message{}
	for i := range s.history[room].Data {
		if strings.Contains(s.history[room].Data[i].Data, val) {
			ret = append(ret, s.history[room].Data[i])
		}
	}
	return ret
}
func (s *SimpleStorage) GlobalSearch(val string) map[string][]types.Message {
	ret := map[string][]types.Message{}
	for k := range s.history {
		for i := range s.history[k].Data {
			if strings.Contains(s.history[k].Data[i].Data, val) {
				ret[k] = append(ret[k], s.history[k].Data[i])
			}
		}
	}
	return ret
}
