package pushname

import (
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type pushname struct {
	pn *events.PushName
}

func New(p *events.PushName) *pushname {
	return &pushname{
		pn: p,
	}
}

func (p *pushname) String() string {
	s := "PushName:"
	s = tools.Setting(s, "JID", tools.JIDString(&p.pn.JID))
	s = tools.Setting(s, "MessageInfo", tools.MessageInfoString(p.pn.Message))
	s = tools.Setting(s, "OldPushName", p.pn.OldPushName)
	s = tools.Setting(s, "NewPushName", p.pn.NewPushName)
	return s
}
