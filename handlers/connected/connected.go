package connected

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
)

type handler struct{}

func init() {
	handlers.Register(handlers.Connected, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	fmt.Println("Connected")
	return nil
}
