package keepaliverestored

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
)

type handler struct{}

func init() {
	handlers.Register(handlers.KeepAliveRestored, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	fmt.Println("KeepAliveRestored")
	return nil
}
