package picture

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type picture struct {
	pic *events.Picture
}

func New(p *events.Picture) *picture {
	return &picture{
		pic: p,
	}
}

func (p *picture) String() string {
	s := "Picture:"
	s = tools.Setting(s, "JID", tools.JIDString(&p.pic.JID))
	s = tools.Setting(s, "Author", tools.JIDString(&p.pic.Author))
	s = tools.Setting(s, "Timestamp", fmt.Sprintf("%v", p.pic.Timestamp))
	s = tools.Setting(s, "Remove", tools.BoolStringIfTrue(p.pic.Remove))
	s = tools.Setting(s, "PictureID", p.pic.PictureID)
	return s
}
