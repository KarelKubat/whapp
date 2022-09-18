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

func MessageInfoString(m *types.MessageInfo) string {
	if m == nil {
		return ""
	}
	s := ""
	s = Setting(s, "ID", m.ID)
	s = Setting(s, "Type", m.Type)
	s = Setting(s, "PushName", m.PushName)
	s = Setting(s, "TimeStamp", fmt.Sprintf("%v", m.Timestamp))
	s = Setting(s, "Category", m.Category)
	s = Setting(s, "Multicast", fmt.Sprintf("%v", m.Multicast))
	s = Setting(s, "MediaType", m.MediaType)
	s += " rest=omitted"
	return s
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

func ExtendedTextMessageString(t *proto.ExtendedTextMessage) string {
	if t == nil {
		return ""
	}
	return fmt.Sprintf("Text=%v, Description=%v, Title=%v, rest=omitted", t.GetText(), t.GetDescription(), t.GetTitle())
}

func DocumentMessageString(d *proto.DocumentMessage) string {
	if d == nil {
		return ""
	}
	return fmt.Sprintf("Title=%v, FileLength=%v, FileName=%v, ContactVcard=%v, rest=omitted", d.GetTitle(), d.GetFileLength(), d.GetFileName(), d.GetContactVcard())
}

func AudioMessageString(a *proto.AudioMessage) string {
	if a == nil {
		return ""
	}
	return fmt.Sprintf("Url=%v, Mimetype=%v, FileLength=%v, Seconds=%v, rest=omitted", a.GetUrl(), a.GetMimetype(), a.GetFileLength(), a.GetSeconds())
}
