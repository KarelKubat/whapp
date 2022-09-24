package archive

import (
	"fmt"

	"github.com/KarelKubat/whapp/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.Archive, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	a := ev.(*events.Archive)
	s := "Archive:"
	s = tools.Setting(s, "JID", tools.JIDString(&a.JID))
	s = tools.Setting(s, "Timestamp", fmt.Sprintf("%v", a.Timestamp))
	s = tools.Setting(s, "Action", fmt.Sprintf("archived=%v,messagerange=%v", a.Action.GetArchived(), a.Action.GetMessageRange()))
	fmt.Println(s)

	return nil
}
