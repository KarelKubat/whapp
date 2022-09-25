// Package message implements the event handler for messages.
package message

import (
	"fmt"

	"github.com/KarelKubat/whapp/tools"
	"github.com/KarelKubat/whatsmeow/handlers"
	"go.mau.fi/whatsmeow/types/events"
)

type handler struct{}

func init() {
	handlers.Register(handlers.Message, &handler{})
}

func (h *handler) Handle(ev interface{}) error {
	m := ev.(*events.Message)
	s := "Message:"

	s = tools.Setting(s, "Info", tools.MessageInfoString(&m.Info))
	s = tools.Setting(s, "Conversation", m.Message.GetConversation())
	s = tools.Setting(s, "Message (unwrapped)", m.Message.String())

	// Unwrapping code incase we need more details.
	// s = tools.Setting(s, "SenderKeyDistributionMessage", tools.SenderKeyDistributionMessageString(m.Message.GetSenderKeyDistributionMessage()))
	// s = tools.Setting(s, "ImageMessage", tools.ImageMessageString(m.Message.GetImageMessage()))
	// s = tools.Setting(s, "ContactMessage", tools.ContactMessageString(m.Message.GetContactMessage()))
	// s = tools.Setting(s, "LocationMessage", tools.LocationMessageString(m.Message.GetLocationMessage()))
	// s = tools.Setting(s, "ExtendedTextMessage", tools.ExtendedTextMessageString(m.Message.ExtendedTextMessage))
	// s = tools.Setting(s, "DocumentMessage", tools.DocumentMessageString(m.Message.DocumentMessage))
	// s = tools.Setting(s, "AudioMessage", tools.AudioMessageString(m.Message.AudioMessage))

	fmt.Println(s)

	return nil
}
