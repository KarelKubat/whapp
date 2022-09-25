package loggedout

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.LoggedOut, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	l := ev.(*events.LoggedOut)
	s := "LoggedOut:"
	s = tools.Setting(s, "OnConnect", fmt.Sprintf("%v", l.OnConnect))
	s = tools.Setting(s, "Reason", fmt.Sprintf("%v", l.Reason))
	fmt.Println(s)

	return nil
}
