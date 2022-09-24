package chatpresence

import (
	"fmt"

	"github.com/KarelKubat/whapp/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.ChatPresence, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	c := ev.(*events.ChatPresence)
	s := "ChatPresence:"
	s = tools.Setting(s, "MessageSource", c.SourceString())
	s = tools.Setting(s, "State", string(c.State))
	s = tools.Setting(s, "Media", string(c.Media))
	fmt.Println(s)

	return nil
}
