package pairerror

import (
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type pairerror struct {
	pe *events.PairError
}

func New(p *events.PairError) *pairerror {
	return &pairerror{
		pe: p,
	}
}

func (p *pairerror) String() string {
	s := "PairError:"
	s = tools.Setting(s, "ID", tools.JIDString(&p.pe.ID))
	s = tools.Setting(s, "BusinessName", p.pe.BusinessName)
	s = tools.Setting(s, "Platform", p.pe.Platform)
	s = tools.Setting(s, "Error", p.pe.Error.Error())
	return s
}
