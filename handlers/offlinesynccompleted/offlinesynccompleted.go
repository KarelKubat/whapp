package offlinesynccompleted

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"github.com/KarelKubat/whatsmeow/handlers"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.OfflineSyncCompleted, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	o := ev.(*events.OfflineSyncCompleted)
	s := "OfflineSyncCompleted:"
	s = tools.Setting(s, "Count", fmt.Sprintf("%v", o.Count))
	fmt.Println(s)

	return nil
}
