package pairerror

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.PairError, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	p := ev.(*events.PairError)
	s := "PairError:"
	s = tools.Setting(s, "ID", tools.JIDString(&p.ID))
	s = tools.Setting(s, "BusinessName", p.BusinessName)
	s = tools.Setting(s, "Platform", p.Platform)
	s = tools.Setting(s, "Error", p.Error.Error())
	fmt.Println(s)

	return nil
}
