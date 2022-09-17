package keepalivetimeout

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type keepalivetimeout struct {
	kt *events.KeepAliveTimeout
}

func New(k *events.KeepAliveTimeout) *keepalivetimeout {
	return &keepalivetimeout{
		kt: k,
	}
}

func (k *keepalivetimeout) String() string {
	s := "KeepAliveTimeout"
	s = tools.Setting(s, "ErrorCount", fmt.Sprintf("%v", k.kt.ErrorCount))
	s = tools.Setting(s, "LastSuccess", fmt.Sprintf("%v", k.kt.LastSuccess))
	return s
}
