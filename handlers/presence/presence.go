package presence

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type presence struct {
	pr *events.Presence
}

func New(p *events.Presence) *presence {
	return &presence{
		pr: p,
	}
}

func (p *presence) String() string {
	s := "Presence:"
	s = tools.Setting(s, "From", tools.JIDString(&p.pr.From))
	s = tools.Setting(s, "Unavailable", fmt.Sprintf("%v", p.pr.Unavailable))
	s = tools.Setting(s, "LastSeen", fmt.Sprintf("%v", p.pr.LastSeen))
	return s
}
