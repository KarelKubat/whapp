package contact

import (
	"fmt"

	"github.com/KarelKubat/whapp/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.Contact, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	c := ev.(*events.Contact)
	s := "Contact:"
	s = tools.Setting(s, "JID", tools.JIDString(&c.JID))
	s = tools.Setting(s, "Timestamp", fmt.Sprintf("%v", c.Timestamp))
	s = tools.Setting(s, "Action", fmt.Sprintf("first:%v, full:%v", c.Action.GetFirstName(), c.Action.GetFullName()))
	fmt.Println(s)

	return nil
}
