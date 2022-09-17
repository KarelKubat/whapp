package loggedout

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type loggedout struct {
	lo *events.LoggedOut
}

func New(l *events.LoggedOut) *loggedout {
	return &loggedout{
		lo: l,
	}
}

func (l *loggedout) String() string {
	s := "LoggedOut:"
	s = tools.Setting(s, "OnConnect", fmt.Sprintf("%v", l.lo.OnConnect))
	s = tools.Setting(s, "Reason", fmt.Sprintf("%v", l.lo.Reason))
	return s
}
