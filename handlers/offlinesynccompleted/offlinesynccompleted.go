package offlinesynccompleted

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type offlinesynccompleted struct {
	osc *events.OfflineSyncCompleted
}

func New(o *events.OfflineSyncCompleted) *offlinesynccompleted {
	return &offlinesynccompleted{
		osc: o,
	}
}

func (o *offlinesynccompleted) String() string {
	s := "OfflineSyncCompleted:"
	s = tools.Setting(s, "Count", fmt.Sprintf("%v", o.osc.Count))
	return s
}
