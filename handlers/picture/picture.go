package picture

import (
	"fmt"

	"github.com/KarelKubat/whapp/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.Picture, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	p := ev.(*events.Picture)
	s := "Picture:"
	s = tools.Setting(s, "JID", tools.JIDString(&p.JID))
	s = tools.Setting(s, "Author", tools.JIDString(&p.Author))
	s = tools.Setting(s, "Timestamp", fmt.Sprintf("%v", p.Timestamp))
	s = tools.Setting(s, "Remove", tools.BoolStringIfTrue(p.Remove))
	s = tools.Setting(s, "PictureID", p.PictureID)
	fmt.Println(s)

	return nil
}
