package keepaliverestored

import "go.mau.fi/whatsmeow/types/events"

type keepaliverestored struct{}

func New(_ *events.KeepAliveRestored) *keepaliverestored {
	return &keepaliverestored{}
}

func (k *keepaliverestored) String() string {
	return "KeepAliveRestored"
}
