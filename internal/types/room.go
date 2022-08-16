package types

type Room struct {
	ID             string
	Name           string
	Participants   []Person
	PinnedMessages []string
	History        MessageHistory
}
