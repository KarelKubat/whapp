package clientoutdated

import "go.mau.fi/whatsmeow/types/events"

type clientoutdated struct{}

func New(c *events.ClientOutdated) *clientoutdated {
	return &clientoutdated{}
}

func (c *clientoutdated) String() string {
	return "ClientOutdated"
}
