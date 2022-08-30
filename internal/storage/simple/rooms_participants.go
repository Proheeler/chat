package simple

import (
	"chat/internal/types"
)

func (s *SimpleStorage) AddParticipantInRoom(patricipant int64, room string) {
	rm := s.rooms[room]
	rm.Participants = append(s.rooms[room].Participants, patricipant)
	part := s.clients[patricipant]
	// rooms := append(part.Rooms, room)
	// part.Rooms = rooms
	s.clients[patricipant] = part
	s.rooms[room] = rm
}
func (s *SimpleStorage) ListParticipantsInRoom(room string) types.ClientList {
	var parts []types.Client
	for i := range s.rooms[room].Participants {
		parts = append(parts, s.clients[s.rooms[room].Participants[i]])
	}
	return types.ClientList{
		Data:  parts,
		Total: len(parts),
	}
}

func (s *SimpleStorage) DeleteParticipantInRoom(uid int64, room string) {
	pl := s.rooms[room].Participants
	rm := s.rooms[room]
	for i := range pl {
		if pl[i] == int64(uid) {
			rm.Participants = RemoveIndexUint(pl, i)
			break
		}
	}
	s.rooms[room] = rm
	cl := s.clients[uid]
	// for i := range cl.Rooms {
	// 	if rm.Name == cl.Rooms[i] {
	// 		cl.Rooms = RemoveIndex(cl.Rooms, i)
	// 		break
	// 	}
	// }
	s.clients[uid] = cl
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func RemoveIndexUint(s []int64, index int) []int64 {
	return append(s[:index], s[index+1:]...)
}
