package connected

import "go.mau.fi/whatsmeow/types/events"

type connected struct{}

func New(c *events.Connected) *connected {
	return &connected{}
}

func (c *connected) String() string {
	return "Connected"
}
