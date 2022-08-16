package logic

import (
	"chat/internal/storage"
	"chat/internal/types"
	"encoding/json"
	"fmt"
	"time"
)

type message struct {
	data []byte
	room string
}

type Subscription struct {
	conn *connection
	room string
	h    *Hub
}

// Hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	// Registered connections.
	rooms map[string]map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan message

	// Register requests from the connections.
	register chan Subscription

	// Unregister requests from connections.
	unregister chan Subscription
	messages   map[string][]string
	storage    storage.Storage
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan message),
		register:   make(chan Subscription),
		unregister: make(chan Subscription),
		rooms:      make(map[string]map[*connection]bool),
		messages:   map[string][]string{},
		storage:    storage.NewSimpleStorage(),
	}
}

func (h *Hub) Storage() storage.Storage {
	return h.storage
}

func (h *Hub) Run() {
	for {
		select {
		case s := <-h.register:
			connections := h.rooms[s.room]
			if connections == nil {
				connections = make(map[*connection]bool)
				h.rooms[s.room] = connections
			}
			h.rooms[s.room][s.conn] = true
		case s := <-h.unregister:
			connections := h.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.room)
					}
				}
			}
		case m := <-h.broadcast:
			connections := h.rooms[m.room]
			for c := range connections {
				select {
				case c.send <- m.data:
					msg := &types.Message{}
					err := json.Unmarshal(m.data, msg)
					if err != nil {
						fmt.Print(err.Error())
						break
					}
					msg.CreationTime = time.Now()
					msg.EditTime = time.Now()
					h.storage.StoreMessage(*msg, m.room)
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, m.room)
					}
				}
			}
		}
	}
}
