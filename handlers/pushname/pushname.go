package pushname

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.PushName, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	p := ev.(*events.PushName)
	s := "PushName:"
	s = tools.Setting(s, "JID", tools.JIDString(&p.JID))
	s = tools.Setting(s, "MessageInfo", tools.MessageInfoString(p.Message))
	s = tools.Setting(s, "OldPushName", p.OldPushName)
	s = tools.Setting(s, "NewPushName", p.NewPushName)
	fmt.Println(s)

	return nil
}
