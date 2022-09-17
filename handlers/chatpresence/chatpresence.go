package chatpresence

import (
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type chatpresence struct {
	cp *events.ChatPresence
}

func New(c *events.ChatPresence) *chatpresence {
	return &chatpresence{
		cp: c,
	}
}

func (c *chatpresence) String() string {
	s := "ChatPresence:"
	s = tools.Setting(s, "MessageSource", c.cp.SourceString())
	s = tools.Setting(s, "State", string(c.cp.State))
	s = tools.Setting(s, "Media", string(c.cp.Media))
	return s
}
