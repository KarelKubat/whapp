package disconnected

import (
	"fmt"

	"github.com/KarelKubat/whapp/handlers"
)

type handler struct{}

func init() {
	handlers.Register(handlers.Disconnected, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	fmt.Println("Disonnected")
	return nil
}
