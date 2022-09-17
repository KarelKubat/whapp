package temporaryban

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type temporaryban struct {
	tb *events.TemporaryBan
}

func New(t *events.TemporaryBan) *temporaryban {
	return &temporaryban{
		tb: t,
	}
}

func (t *temporaryban) String() string {
	s := "TemporaryBan:"
	s = tools.Setting(s, "Code", t.tb.Code.String())
	s = tools.Setting(s, "Expire", fmt.Sprintf("%v", t.tb.Expire))
	return s
}
