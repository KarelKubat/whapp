package temporaryban

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.TemporaryBan, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	t := ev.(*events.TemporaryBan)
	s := "TemporaryBan:"
	s = tools.Setting(s, "Code", t.Code.String())
	s = tools.Setting(s, "Expire", fmt.Sprintf("%v", t.Expire))
	fmt.Println(s)

	return nil
}
