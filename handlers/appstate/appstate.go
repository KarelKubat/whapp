package appstate

import "go.mau.fi/whatsmeow/types/events"

type appstate struct {
	as *events.AppState
}

func New(a *events.AppState) *appstate {
	return &appstate{
		as: a,
	}
}

func (a *appstate) String() string {
	return "AppState: omitted"
}
