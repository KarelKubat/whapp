package appstatesynccomplete

import "go.mau.fi/whatsmeow/types/events"

type appstatesynccomplete struct {
	as *events.AppStateSyncComplete
}

func New(a *events.AppStateSyncComplete) *appstatesynccomplete {
	return &appstatesynccomplete{
		as: a,
	}
}

func (a *appstatesynccomplete) String() string {
	return "AppStateSyncComplete: omitted"
}
