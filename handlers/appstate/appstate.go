package appstate

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
)

type handler struct{}

func init() {
	handlers.Register(handlers.AppState, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	fmt.Println("AppState: [omitted]")
	return nil
}
