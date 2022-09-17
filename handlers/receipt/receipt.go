package receipt

import (
	"fmt"
	"strings"

	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type receipt struct {
	rec *events.Receipt
}

func New(r *events.Receipt) *receipt {
	return &receipt{
		rec: r,
	}
}

func (r *receipt) String() string {
	s := "Receipt:"
	s = tools.Setting(s, "MessageSource", r.rec.MessageSource.SourceString())
	s = tools.Setting(s, "MessageIDs", strings.Join(r.rec.MessageIDs, ","))
	s = tools.Setting(s, "Timestamp", fmt.Sprintf("%v", r.rec.Timestamp))
	s = tools.Setting(s, "Type", string(r.rec.Type))
	return s
}
