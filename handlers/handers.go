package handlers

import (
	"fmt"

	"github.com/KarelKubat/whapp/handlers/appstate"
	"github.com/KarelKubat/whapp/handlers/appstatesynccomplete"
	"github.com/KarelKubat/whapp/handlers/archive"
	"github.com/KarelKubat/whapp/handlers/chatpresence"
	"github.com/KarelKubat/whapp/handlers/clientoutdated"
	"github.com/KarelKubat/whapp/handlers/connected"
	"github.com/KarelKubat/whapp/handlers/contact"
	"github.com/KarelKubat/whapp/handlers/disconnected"
	"github.com/KarelKubat/whapp/handlers/keepaliverestored"
	"github.com/KarelKubat/whapp/handlers/keepalivetimeout"
	"github.com/KarelKubat/whapp/handlers/loggedout"
	"github.com/KarelKubat/whapp/handlers/mediaretryerror"
	"github.com/KarelKubat/whapp/handlers/message"
	"github.com/KarelKubat/whapp/handlers/offlinesynccompleted"
	"github.com/KarelKubat/whapp/handlers/pairsuccess"
	"github.com/KarelKubat/whapp/handlers/picture"
	"github.com/KarelKubat/whapp/handlers/presence"
	"github.com/KarelKubat/whapp/handlers/privacysettings"
	"github.com/KarelKubat/whapp/handlers/qr"
	"github.com/KarelKubat/whapp/handlers/receipt"
	"github.com/KarelKubat/whapp/handlers/streamreplaced"
	"github.com/KarelKubat/whapp/handlers/temporaryban"
	"github.com/KarelKubat/whapp/handlers/unknown"

	"go.mau.fi/whatsmeow/types/events"
)

func Dispatch(evt interface{}) {
	switch v := evt.(type) {
	case *events.AppState:
		fmt.Println(appstate.New(v).String())
	case *events.AppStateSyncComplete:
		fmt.Println(appstatesynccomplete.New(v).String())
	case *events.Archive:
		fmt.Println(archive.New(v).String())
	case *events.ChatPresence:
		fmt.Println(chatpresence.New(v).String())
	case *events.ClientOutdated:
		fmt.Println(clientoutdated.New(v).String())
	case *events.Contact:
		fmt.Println(contact.New(v).String())
	case *events.Connected:
		fmt.Println(connected.New(v).String())
	case *events.Disconnected:
		fmt.Println(disconnected.New(v).String())
	case *events.KeepAliveRestored:
		fmt.Println(keepaliverestored.New(v).String())
	case *events.KeepAliveTimeout:
		fmt.Println(keepalivetimeout.New(v).String())
	case *events.LoggedOut:
		fmt.Println(loggedout.New(v).String())
	case *events.MediaRetryError:
		fmt.Println(mediaretryerror.New(v).String())
	case *events.Message:
		fmt.Println(message.New(v).String())
	case *events.OfflineSyncCompleted:
		fmt.Println(offlinesynccompleted.New(v).String())
	case *events.PairSuccess:
		fmt.Println(pairsuccess.New(v).String())
	case *events.Picture:
		fmt.Println(picture.New(v).String())
	case *events.PrivacySettings:
		fmt.Println(privacysettings.New(v).String())
	case *events.Presence:
		fmt.Println(presence.New(v).String())
	case *events.QR:
		fmt.Println(qr.New(v).String())
	case *events.Receipt:
		fmt.Println(receipt.New(v).String())
	case *events.StreamReplaced:
		fmt.Println(streamreplaced.New(v).String())
	case *events.TemporaryBan:
		fmt.Println(temporaryban.New(v).String())
	default:
		fmt.Println(unknown.New(fmt.Sprintf("%+v", evt)))
	}
}
