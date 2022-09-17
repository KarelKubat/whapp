package pairsuccess

import (
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type pairsuccess struct {
	ps *events.PairSuccess
}

func New(p *events.PairSuccess) *pairsuccess {
	return &pairsuccess{
		ps: p,
	}
}

func (p *pairsuccess) String() string {
	s := "PairSuccess:"
	s = tools.Setting(s, "ID", tools.JIDString(&p.ps.ID))
	s = tools.Setting(s, "BusinessName", p.ps.BusinessName)
	s = tools.Setting(s, "Platform", p.ps.Platform)
	return s
}
