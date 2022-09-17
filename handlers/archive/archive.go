package archive

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type archive struct {
	arc *events.Archive
}

func New(a *events.Archive) *archive {
	return &archive{
		arc: a,
	}
}

func (a *archive) String() string {
	s := "Archive:"
	s = tools.Setting(s, "JID", tools.JIDString(&a.arc.JID))
	s = tools.Setting(s, "Timestamp", fmt.Sprintf("%v", a.arc.Timestamp))
	s = tools.Setting(s, "Action", fmt.Sprintf("archived=%v,messagerange=%v", a.arc.Action.GetArchived(), a.arc.Action.GetMessageRange()))

	return s
}
