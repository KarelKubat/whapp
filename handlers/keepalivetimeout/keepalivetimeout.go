package keepalivetimeout

import (
	"fmt"

	"github.com/KarelKubat/whapp/handlers"
)

type handler struct{}

func init() {
	handlers.Register(handlers.KeepAliveTimeout, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	fmt.Println("KeepAliveTimeout")
	return nil
}
