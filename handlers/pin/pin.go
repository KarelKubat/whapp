package pin

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type pin struct {
	pn *events.Pin
}

func New(p *events.Pin) *pin {
	return &pin{
		pn: p,
	}
}

func (p *pin) String() string {
	s := "Pin:"
	s = tools.Setting(s, "JID", tools.JIDString(&p.pn.JID))
	s = tools.Setting(s, "Timestamp", fmt.Sprintf("%v", p.pn.Timestamp))
	s = tools.Setting(s, "Action,pinned", fmt.Sprintf("%v", p.pn.Action.GetPinned()))
	return s
}
