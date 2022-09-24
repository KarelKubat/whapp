package handlers

import (
	"fmt"
	"sync"

	"go.mau.fi/whatsmeow/types/events"
)

type Type int

const (
	firstUnused Type = iota // keep at first slot for tests

	AppState
	AppStateSyncComplete
	Archive
	BusinessName
	CallAccept
	CallOffer
	CallOfferNotice
	CallRelayLatency
	CallTerminate
	ChatPresence
	ClientOutdated
	Connected
	ConnectFailure
	Contact
	DeleteChat
	DeleteForMe
	Disconnected
	GroupInfo
	HistorySync
	IdentityChange
	JoinedGroup
	KeepAliveRestored
	KeepAliveTimeout
	LoggedOut
	MarkChatAsRead
	MediaRetry
	Message
	Mute
	OfflineSyncCompleted
	OfflineSyncPreview
	PairError
	PairSuccess
	Picture
	Pin
	Presence
	PrivacySettings
	PushName
	PushNameSetting
	QR
	QRScannedWithoutMultidevice
	Receipt
	Star
	StreamError
	StreamReplaced
	TemporaryBan
	UnarchiveChatSetting
	UndecryptableMessage
	UnknownCallEvent

	lastUnused // keep at last slot for tests
)

func (t Type) String() string {
	return []string{
		"", // unused
		"AppState",
		"AppStateSyncComplete",
		"Archive",
		"BusinessName",
		"CallAccept",
		"CallOffer",
		"CallOfferNotice",
		"CallRelayLatency",
		"CallTerminate",
		"ChatPresence",
		"ClientOutdated",
		"Connected",
		"ConnectFailure",
		"Contact",
		"DeleteChat",
		"DeleteForMe",
		"Disconnected",
		"GroupInfo",
		"HistorySync",
		"IdentityChange",
		"JoinedGroup",
		"KeepAliveRestored",
		"KeepAliveTimeout",
		"LoggedOut",
		"MarkChatAsRead",
		"MediaRetry",
		"Message",
		"Mute",
		"OfflineSyncCompleted",
		"OfflineSyncPreview",
		"PairError",
		"PairSuccess",
		"Picture",
		"Pin",
		"Presence",
		"PrivacySettings",
		"PushName",
		"PushNameSetting",
		"QR",
		"QRScannedWithoutMultidevice",
		"Receipt",
		"Star",
		"StreamError",
		"StreamReplaced",
		"TemporaryBan",
		"UnarchiveChatSetting",
		"UndecryptableMessage",
		"UnknownCallEvent",
	}[t]
}

type handler interface {
	Handle(evt interface{}) error
}

var registry = make(map[Type][]handler)
var registryMutex sync.Mutex

func Register(t Type, h handler) {
	registryMutex.Lock()
	defer registryMutex.Unlock()

	if _, ok := registry[t]; !ok {
		registry[t] = []handler{}
	}
	registry[t] = append(registry[t], h)
}

func Dispatch(evt interface{}) error {
	switch v := evt.(type) {
	case *events.AppState:
		return dispatch(AppState, v)
	case *events.AppStateSyncComplete:
		return dispatch(AppStateSyncComplete, v)
	case *events.Archive:
		return dispatch(Archive, v)
	case *events.BusinessName:
		return dispatch(BusinessName, v)
	case *events.CallAccept:
		return dispatch(CallAccept, v)
	case *events.CallOffer:
		return dispatch(CallOffer, v)
	case *events.CallOfferNotice:
		return dispatch(CallOfferNotice, v)
	case *events.CallRelayLatency:
		return dispatch(CallRelayLatency, v)
	case *events.CallTerminate:
		return dispatch(CallTerminate, v)
	case *events.ChatPresence:
		return dispatch(ChatPresence, v)
	case *events.ClientOutdated:
		return dispatch(ClientOutdated, v)
	case *events.Connected:
		return dispatch(Connected, v)
	case *events.ConnectFailure:
		return dispatch(ConnectFailure, v)
	case *events.Contact:
		return dispatch(Contact, v)
	case *events.DeleteChat:
		return dispatch(DeleteChat, v)
	case *events.DeleteForMe:
		return dispatch(DeleteForMe, v)
	case *events.Disconnected:
		return dispatch(Disconnected, v)
	case *events.GroupInfo:
		return dispatch(GroupInfo, v)
	case *events.HistorySync:
		return dispatch(HistorySync, v)
	case *events.JoinedGroup:
		return dispatch(JoinedGroup, v)
	case *events.IdentityChange:
		return dispatch(IdentityChange, v)
	case *events.KeepAliveRestored:
		return dispatch(KeepAliveRestored, v)
	case *events.KeepAliveTimeout:
		return dispatch(KeepAliveTimeout, v)
	case *events.LoggedOut:
		return dispatch(LoggedOut, v)
	case *events.MarkChatAsRead:
		return dispatch(MarkChatAsRead, v)
	case *events.MediaRetry:
		return dispatch(MediaRetry, v)
	case *events.Message:
		return dispatch(Message, v)
	case *events.OfflineSyncCompleted:
		return dispatch(OfflineSyncCompleted, v)
	case *events.OfflineSyncPreview:
		return dispatch(OfflineSyncPreview, v)
	case *events.PairError:
		return dispatch(PairError, v)
	case *events.PairSuccess:
		return dispatch(PairSuccess, v)
	case *events.Picture:
		return dispatch(Picture, v)
	case *events.Pin:
		return dispatch(Pin, v)
	case *events.Presence:
		return dispatch(Presence, v)
	case *events.PrivacySettings:
		return dispatch(PrivacySettings, v)
	case *events.PushName:
		return dispatch(PushName, v)
	case *events.PushNameSetting:
		return dispatch(PushNameSetting, v)
	case *events.QR:
		return dispatch(QR, v)
	case *events.QRScannedWithoutMultidevice:
		return dispatch(QRScannedWithoutMultidevice, v)
	case *events.Receipt:
		return dispatch(Receipt, v)
	case *events.Star:
		return dispatch(Star, v)
	case *events.StreamError:
		return dispatch(StreamError, v)
	case *events.StreamReplaced:
		return dispatch(StreamReplaced, v)
	case *events.TemporaryBan:
		return dispatch(TemporaryBan, v)
	case *events.UnarchiveChatsSetting:
		return dispatch(UnarchiveChatSetting, v)
	case *events.UndecryptableMessage:
		return dispatch(UndecryptableMessage, v)
	case *events.UnknownCallEvent:
		return dispatch(UnknownCallEvent, v)
	default:
		return fmt.Errorf("unknown event %+v, can't dispatch", v)
	}
}

func dispatch(t Type, ev interface{}) error {
	if handlers, ok := registry[t]; ok {
		for _, h := range handlers {
			if err := h.Handle(ev); err != nil {
				return err
			}
		}
		return nil
	}
	return fmt.Errorf("%v: no handler for event %+v", t, ev)
}
