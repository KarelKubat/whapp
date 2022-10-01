package poll

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KarelKubat/whapp/action"
	"github.com/KarelKubat/whapp/wm"

	"github.com/KarelKubat/whapp/flags"
	"github.com/KarelKubat/whatsmeow/handlers"

	_ "github.com/KarelKubat/whapp/handlers/appstate"
	_ "github.com/KarelKubat/whapp/handlers/appstatesynccomplete"
	_ "github.com/KarelKubat/whapp/handlers/archive"
	_ "github.com/KarelKubat/whapp/handlers/chatpresence"
	_ "github.com/KarelKubat/whapp/handlers/clientoutdated"
	_ "github.com/KarelKubat/whapp/handlers/connected"
	_ "github.com/KarelKubat/whapp/handlers/contact"
	_ "github.com/KarelKubat/whapp/handlers/deletechat"
	_ "github.com/KarelKubat/whapp/handlers/disconnected"
	_ "github.com/KarelKubat/whapp/handlers/identitychange"
	_ "github.com/KarelKubat/whapp/handlers/keepaliverestored"
	_ "github.com/KarelKubat/whapp/handlers/keepalivetimeout"
	_ "github.com/KarelKubat/whapp/handlers/loggedout"
	_ "github.com/KarelKubat/whapp/handlers/message"
	_ "github.com/KarelKubat/whapp/handlers/offlinesynccompleted"
	_ "github.com/KarelKubat/whapp/handlers/offlinesyncpreview"
	_ "github.com/KarelKubat/whapp/handlers/pairerror"
	_ "github.com/KarelKubat/whapp/handlers/pairsuccess"
	_ "github.com/KarelKubat/whapp/handlers/picture"
	_ "github.com/KarelKubat/whapp/handlers/pin"
	_ "github.com/KarelKubat/whapp/handlers/presence"
	_ "github.com/KarelKubat/whapp/handlers/privacysettings"
	_ "github.com/KarelKubat/whapp/handlers/pushname"
	_ "github.com/KarelKubat/whapp/handlers/qr"
	_ "github.com/KarelKubat/whapp/handlers/receipt"
	_ "github.com/KarelKubat/whapp/handlers/streamreplaced"
	_ "github.com/KarelKubat/whapp/handlers/temporaryban"
	_ "github.com/KarelKubat/whapp/handlers/undecryptablemessage"
)

type poll struct{}

func init() {
	action.Register(&poll{})
}

func (p *poll) Description() *action.Info {
	return &action.Info{
		Name:  "poll",
		NArgs: 0,
		Usage: "poll (continuously listen and show events on stdout until -quit-after runs out or ^C is pressed)",
	}
}

func (p *poll) Run(_ []string) error {
	// Instantiate client.
	if err := wm.InitClient(); err != nil {
		return err
	}

	// The handlers have registered themselves. Make sure to dispatch seen events.
	wm.Client.AddEventHandler(func(e interface{}) {
		if err := handlers.Dispatch(e); err != nil {
			switch err.Type {
			case handlers.NoHandlerFound:
				// This program tries to cover all events.
				fmt.Println(err)
			case handlers.HandlerFailed:
				// Handler has failed.
				fmt.Println(err)
				// Dispatcher has a bug or lags in known events.
			case handlers.UnknownEvent:
				panic(err)
			}
		}
	})

	// Poll forever, or stop after a tiven time.
	if *flags.QuitAfterFlag == 0 {
		fmt.Println("Press ^C to stop.")
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
	} else {
		fmt.Println("Press ^C to stop, or wait for", *flags.QuitAfterFlag)
		time.Sleep(*flags.QuitAfterFlag)
	}

	// Cleanup.
	wm.Client.Disconnect()
	return nil
}
