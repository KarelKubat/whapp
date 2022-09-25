package presence

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.Presence, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	p := ev.(*events.Presence)
	s := "Presence:"
	s = tools.Setting(s, "From", tools.JIDString(&p.From))
	s = tools.Setting(s, "Unavailable", fmt.Sprintf("%v", p.Unavailable))
	s = tools.Setting(s, "LastSeen", fmt.Sprintf("%v", p.LastSeen))
	fmt.Println(s)

	return nil
}
