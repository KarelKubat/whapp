package pairsuccess

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.PairSuccess, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	p := ev.(*events.PairSuccess)
	s := "PairSuccess:"
	s = tools.Setting(s, "ID", tools.JIDString(&p.ID))
	s = tools.Setting(s, "BusinessName", p.BusinessName)
	s = tools.Setting(s, "Platform", p.Platform)
	fmt.Println(s)

	return nil
}
