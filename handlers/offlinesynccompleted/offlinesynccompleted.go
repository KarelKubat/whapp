package offlinesynccompleted

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.Message, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	o := ev.(*events.OfflineSyncCompleted)
	s := "OfflineSyncCompleted:"
	s = tools.Setting(s, "Count", fmt.Sprintf("%v", o.Count))
	fmt.Println(s)

	return nil
}
