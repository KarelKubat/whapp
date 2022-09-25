// Package message implements the event handler for messages.
package message

import (
	"fmt"

	"github.com/KarelKubat/whatsmeow/handlers"
	"github.com/KarelKubat/whapp/tools"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.Message, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	m := ev.(*events.Message)
	s := "Message:"

	s = tools.Setting(s, "ID", m.Info.ID)
	s = tools.Setting(s, "Type", m.Info.Type)
	s = tools.Setting(s, "PushName", m.Info.PushName)
	s = tools.Setting(s, "Category", m.Info.Category)
	s = tools.Setting(s, "Multicast", tools.BoolStringIfTrue(m.Info.Multicast))
	s = tools.Setting(s, "MediaType", m.Info.MediaType)
	// VerifiedName and DeviceSentData are not rendered

	s = tools.Setting(s, "Conversation", m.Message.GetConversation())

	org := s

	s = tools.Setting(s, "SenderKeyDistributionMessage", tools.SenderKeyDistributionMessageString(m.Message.GetSenderKeyDistributionMessage()))
	s = tools.Setting(s, "ImageMessage", tools.ImageMessageString(m.Message.GetImageMessage()))
	s = tools.Setting(s, "ContactMessage", tools.ContactMessageString(m.Message.GetContactMessage()))
	s = tools.Setting(s, "LocationMessage", tools.LocationMessageString(m.Message.GetLocationMessage()))
	s = tools.Setting(s, "ExtendedTextMessage", tools.ExtendedTextMessageString(m.Message.ExtendedTextMessage))
	s = tools.Setting(s, "DocumentMessage", tools.DocumentMessageString(m.Message.DocumentMessage))
	s = tools.Setting(s, "AudioMessage", tools.AudioMessageString(m.Message.AudioMessage))

	if org == s {
		s += "SUBMESSAGE NOT DECODED, FIX hanlders/message/message.go"
	}

	fmt.Println(s)

	return nil
}
