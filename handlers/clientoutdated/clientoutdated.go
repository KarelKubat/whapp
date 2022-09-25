package clientoutdated

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
)

type handler struct{}

func init() {
	handlers.Register(handlers.ClientOutdated, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	fmt.Println("ClientOutdated")
	return nil
}
