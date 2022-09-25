package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KarelKubat/flagnames"
	"github.com/KarelKubat/whatsmeow/handlers"
	"github.com/KarelKubat/whatsmeow/logger"
	"github.com/mdp/qrterminal/v3"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"

	_ "github.com/mattn/go-sqlite3"

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

var (
	logfileFlag   = flag.String("logfile", "/tmp/whapp.log", "logfile to write")
	verboseFlag   = flag.Bool("verbose", false, "when true, debug messages are logged")
	appendFlag    = flag.Bool("append", true, "when true (default), the logfile is appended")
	dbFlag        = flag.String("db", "store.db", "sqlite3 backend")
	quitAfterFlag = flag.Duration("quit-after", 0, "stop after the given duration (default: poll forever")
)

const (
	usageInfo = `
Usage: whapp [FLAGS]
So far whapp only listens to incoming messages and displays them as they are received.
Accepted flags are listed below. They can be abbreviated (e.g., -v for -verbose).

`
)

func main() {
	// Commandline
	flagnames.Patch()
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 0 {
		usage()
	}

	// Logging
	baseLogger, err := logger.New(logger.Opts{
		Filename: *logfileFlag,
		Verbose:  *verboseFlag,
		Append:   *appendFlag,
	})
	checkErr(err)
	dbLogger := baseLogger.Sub("Database")
	clLogger := baseLogger.Sub("Client")

	// Instantiate storage.
	container, err := sqlstore.New("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=on", *dbFlag), dbLogger)
	checkErr(err)
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	store, err := container.GetFirstDevice()
	checkErr(err)

	// Instantiate client.
	client := whatsmeow.NewClient(store, clLogger)
	client.AddEventHandler(func(e interface{}) {
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

	// Authentication.
	if client.Store.ID == nil {
		// No ID stored, this may be a new auth.
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		checkErr(err)

		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("Scan the QR code to authenticate whapp.")
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err = client.Connect()
		checkErr(err)
	}

	if *quitAfterFlag == 0 {
		fmt.Println("Press ^C to stop.")
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
	} else {
		fmt.Println("Press ^C to stop, or wait for", *quitAfterFlag)
		time.Sleep(*quitAfterFlag)
	}

	client.Disconnect()
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: %v\n", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, usageInfo)
	flag.PrintDefaults()
	fmt.Fprintln(os.Stdout)
}
