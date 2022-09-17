package disconnected

import "go.mau.fi/whatsmeow/types/events"

type disconnected struct{}

func New(c *events.Disconnected) *disconnected {
	return &disconnected{}
}

func (c *disconnected) String() string {
	return "Disconnected"
}
