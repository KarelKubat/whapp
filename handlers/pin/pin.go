package pin

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.Pin, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	p := ev.(*events.Pin)
	s := "Pin:"
	s = tools.Setting(s, "JID", tools.JIDString(&p.JID))
	s = tools.Setting(s, "Timestamp", fmt.Sprintf("%v", p.Timestamp))
	s = tools.Setting(s, "Action,pinned", fmt.Sprintf("%v", p.Action.GetPinned()))
	fmt.Println(s)

	return nil
}
