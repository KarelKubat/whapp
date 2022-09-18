// Package message implements the event handler for messages.
package message

import (
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type message struct {
	ms *events.Message
}

func New(m *events.Message) *message {
	return &message{
		ms: m,
	}
}

func (m *message) String() string {
	s := "Message:"

	s = tools.Setting(s, "ID", m.ms.Info.ID)
	s = tools.Setting(s, "Type", m.ms.Info.Type)
	s = tools.Setting(s, "PushName", m.ms.Info.PushName)
	s = tools.Setting(s, "Category", m.ms.Info.Category)
	s = tools.Setting(s, "Multicast", tools.BoolStringIfTrue(m.ms.Info.Multicast))
	s = tools.Setting(s, "MediaType", m.ms.Info.MediaType)
	// VerifiedName and DeviceSentData are not rendered

	s = tools.Setting(s, "Conversation", m.ms.Message.GetConversation())

	org := s

	s = tools.Setting(s, "SenderKeyDistributionMessage", tools.SenderKeyDistributionMessageString(m.ms.Message.GetSenderKeyDistributionMessage()))
	s = tools.Setting(s, "ImageMessage", tools.ImageMessageString(m.ms.Message.GetImageMessage()))
	s = tools.Setting(s, "ContactMessage", tools.ContactMessageString(m.ms.Message.GetContactMessage()))
	s = tools.Setting(s, "LocationMessage", tools.LocationMessageString(m.ms.Message.GetLocationMessage()))
	s = tools.Setting(s, "ExtendedTextMessage", tools.ExtendedTextMessageString(m.ms.Message.ExtendedTextMessage))
	s = tools.Setting(s, "DocumentMessage", tools.DocumentMessageString(m.ms.Message.DocumentMessage))
	s = tools.Setting(s, "AudioMessage", tools.AudioMessageString(m.ms.Message.AudioMessage))

	if org == s {
		s += "SUBMESSAGE NOT DECODED, ADAPT MESSAGE.GO"
	}

	return s
}
