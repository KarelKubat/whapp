package qr

import (
	"fmt"
	"strings"

	"github.com/KarelKubat/whapp/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.QR, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	q := ev.(*events.QR)
	s := "QR:"
	s = tools.Setting(s, "Codes", strings.Join(q.Codes, " "))
	fmt.Println(s)

	return nil
}
