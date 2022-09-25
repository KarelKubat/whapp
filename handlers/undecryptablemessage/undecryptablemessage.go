package undecryptablemessage

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"github.com/KarelKubat/whatsmeow/handlers"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.UndecryptableMessage, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	u := ev.(*events.UndecryptableMessage)
	s := "UndecryptableMessage:"
	s = tools.Setting(s, "Info", tools.MessageInfoString(&u.Info))
	s = tools.Setting(s, "IsUnavailable", tools.BoolStringIfTrue(u.IsUnavailable))

	fmt.Println(s)

	return nil
}
