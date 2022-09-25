package streamreplaced

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
)

type handler struct{}

func init() {
	handlers.Register(handlers.StreamReplaced, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	fmt.Println("StreamReplaced")
	return nil
}
