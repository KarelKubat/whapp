package appstatesynccomplete

import (
	"fmt"

	"github.com/KarelKubat/whapp/handlers"
)

type handler struct{}

func init() {
	handlers.Register(handlers.AppStateSyncComplete, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	fmt.Println("AppStateSyncComplete: [omitted]")
	return nil
}
