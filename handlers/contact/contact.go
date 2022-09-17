package contact

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type contact struct {
	cn *events.Contact
}

func New(c *events.Contact) *contact {
	return &contact{
		cn: c,
	}
}

func (c *contact) String() string {
	s := "Contact:"
	s = tools.Setting(s, "JID", tools.JIDString(&c.cn.JID))
	s = tools.Setting(s, "Timestamp", fmt.Sprintf("%v", c.cn.Timestamp))
	s = tools.Setting(s, "Action", fmt.Sprintf("first:%v, full:%v", c.cn.Action.GetFirstName(), c.cn.Action.GetFullName()))
	return s
}
