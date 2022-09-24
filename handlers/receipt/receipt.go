package receipt

import (
	"fmt"
	"strings"

	"github.com/KarelKubat/whapp/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.Receipt, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	r := ev.(*events.Receipt)
	s := "Receipt:"
	s = tools.Setting(s, "MessageSource", r.MessageSource.SourceString())
	s = tools.Setting(s, "MessageIDs", strings.Join(r.MessageIDs, ","))
	s = tools.Setting(s, "Timestamp", fmt.Sprintf("%v", r.Timestamp))
	s = tools.Setting(s, "Type", string(r.Type))
	fmt.Println(s)

	return nil
}
