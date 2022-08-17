package simple

import "strings"

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
