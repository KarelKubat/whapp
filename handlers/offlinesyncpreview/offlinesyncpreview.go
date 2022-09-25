package offlinesyncpreview

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"github.com/KarelKubat/whatsmeow/handlers"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.OfflineSyncPreview, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	o := ev.(*events.OfflineSyncPreview)
	s := "OfflineSyncPreview"

	s = tools.Setting(s, "Total", fmt.Sprintf("%v", o.Total))
	s = tools.Setting(s, "AppDataChanges", fmt.Sprintf("%v", o.AppDataChanges))
	s = tools.Setting(s, "Messages", fmt.Sprintf("%v", o.Messages))
	s = tools.Setting(s, "Notifications", fmt.Sprintf("%v", o.Notifications))
	s = tools.Setting(s, "Receipts", fmt.Sprintf("%v", o.Receipts))

	fmt.Println(s)

	return nil
}
