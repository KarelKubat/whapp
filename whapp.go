package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/KarelKubat/flagnames"
	"github.com/KarelKubat/whapp/logger"
	"github.com/mdp/qrterminal/v3"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"

	_ "github.com/mattn/go-sqlite3"
)

var (
	logfileFlag = flag.String("logfile", "/tmp/whapp.log", "logfile to write")
	verboseFlag = flag.Bool("verbose", false, "when true, debug messages are logged")
	appendFlag  = flag.Bool("append", true, "when true (default), the logfile is appended")
	reopenFlag  = flag.Bool("reopen", false, "when true, a new file is reopened when the logfile disappears (useful for log rotation)")
	dbFlag      = flag.String("db", "store.db", "sqlite3 backend")
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
		Reopen:   *reopenFlag,
	})
	checkErr(err)
	dbLogger := baseLogger.Sub("Database")
	clientLogger := baseLogger.Sub("Client")

	// Instantiate storage.
	container, err := sqlstore.New("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=on", *dbFlag), dbLogger)
	checkErr(err)
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	checkErr(err)

	// Instantiate client.
	client := whatsmeow.NewClient(deviceStore, clientLogger)
	client.AddEventHandler(eventHandler)

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

	fmt.Println("Press ^C to stop.")
	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: %v\n", err)
		os.Exit(1)
	}
}

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		if m := v.Message.GetConversation(); m != "" {
			fmt.Println("Received a message:", v.Message.GetConversation())
		} else {
			fmt.Println("Received an empty message")
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, usageInfo)
	flag.PrintDefaults()
	fmt.Fprintln(os.Stdout)
}