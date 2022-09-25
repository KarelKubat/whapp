package deletechat

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"github.com/KarelKubat/whatsmeow/handlers"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.DeleteChat, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	d := ev.(*events.DeleteChat)
	s := "DeleteChat:"

	s = tools.Setting(s, "JID", tools.JIDString(&d.JID))
	s = tools.Setting(s, "Timestamp", fmt.Sprintf("%v", d.Timestamp))
	s = tools.Setting(s, "DeleteChatAction", d.Action.String())

	fmt.Println(s)

	return nil
}
