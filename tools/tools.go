package tools

import (
	"fmt"

	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

func Setting(s, about, payload string) string {
	if payload == "" {
		return s
	}
	if s != "" {
		s += " "
	}
	s += fmt.Sprintf("[%s:%v]", about, payload)
	return s
}

func JIDString(j *types.JID) string {
	s := ""
	s = Setting(s, "User", j.User)
	if j.Agent != 0 {
		s = Setting(s, "Agent", fmt.Sprintf("%v", j.Agent))
	}
	if j.Device != 0 {
		s = Setting(s, "Device", fmt.Sprintf("%v", j.Device))
	}
	s = Setting(s, "Server", j.Server)
	if j.AD {
		s = Setting(s, "AD", "true")
	}
	return s
}

func BoolStringIfTrue(b bool) string {
	if b {
		return "true"
	}
	return ""
}

// Protos (mainly in a Message type)
func SenderKeyDistributionMessageString(s *proto.SenderKeyDistributionMessage) string {
	if s == nil {
		return ""
	}
	return fmt.Sprintf("GroupId=%v, rest=omitted", s.GetGroupId())
}

func ImageMessageString(i *proto.ImageMessage) string {
	if i == nil {
		return ""
	}
	return fmt.Sprintf("Caption=%v, Mimetype=%v, rest=omitted", i.GetCaption(), i.GetMimetype())
}

func ContactMessageString(c *proto.ContactMessage) string {
	if c == nil {
		return ""
	}
	return fmt.Sprintf("DisplayName=%v, rest=omitted", c.GetDisplayName())
}

func LocationMessageString(l *proto.LocationMessage) string {
	if l == nil {
		return ""
	}
	return fmt.Sprintf("Name=%v, rest=omitted", l.GetName())
}
