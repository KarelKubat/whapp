package qr

import (
	"strings"

	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type qr struct {
	codes []string
}

func New(c *events.QR) *qr {
	return &qr{
		codes: c.Codes,
	}
}

func (q *qr) String() string {
	s := "QR:"
	s = tools.Setting(s, "Codes", strings.Join(q.codes, " "))
	return s
}
