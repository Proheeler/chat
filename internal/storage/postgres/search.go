package postgres

import "strings"

func (s *PostgresStorage) Search(val, room string) []int {
	ret := []int{}
	for i := range s.history[room].Data {
		if strings.Contains(s.history[room].Data[i].Data, val) {
			ret = append(ret, i)
		}
	}
	return ret
}
func (s *PostgresStorage) GlobalSearch(val string) map[string][]int {
	ret := map[string][]int{}
	for k := range s.history {
		for i := range s.history[k].Data {
			if strings.Contains(s.history[k].Data[i].Data, val) {
				ret[k] = append(ret[k], i)
			}
		}
	}
	return ret
}
