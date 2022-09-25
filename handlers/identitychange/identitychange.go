package identitychange

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"github.com/KarelKubat/whatsmeow/handlers"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.IdentityChange, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	i := ev.(*events.IdentityChange)
	s := "IdentityChange:"

	s = tools.Setting(s, "JID", tools.JIDString(&i.JID))
	s = tools.Setting(s, "Timestamp", fmt.Sprintf("%v", i.Timestamp))
	s = tools.Setting(s, "Implicit", tools.BoolStringIfTrue(i.Implicit))

	fmt.Println(s)

	return nil
}
