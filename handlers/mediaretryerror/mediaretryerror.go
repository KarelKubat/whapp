package mediaretryerror

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type mediaretryerror struct {
	mr *events.MediaRetryError
}

func New(m *events.MediaRetryError) *mediaretryerror {
	return &mediaretryerror{
		mr: m,
	}
}

func (m *mediaretryerror) String() string {
	s := "MediaRetryError:"
	s = tools.Setting(s, "Code", fmt.Sprintf("%v", m.mr.Code))
	return s
}
