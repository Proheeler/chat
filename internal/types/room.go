package types

type Room struct {
	ID             string
	Name           string
	Participants   []Client
	PinnedMessages []string
	History        *MessageHistory
}
