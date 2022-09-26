package wm

import (
	"context"
	"fmt"
	"os"

	"github.com/KarelKubat/whapp/flags"
	"github.com/KarelKubat/whapp/loggers"
	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"

	_ "github.com/mattn/go-sqlite3"
)

var (
	clientInitialized bool
	Client            *whatsmeow.Client
	storeInitialized  bool
	Store             *store.Device
)

func InitStore() error {
	if storeInitialized {
		return nil
	}

	container, err := sqlstore.New("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=on", *flags.DBFlag), loggers.DBLogger)
	if err != nil {
		return err
	}
	Store, err = container.GetFirstDevice()
	if err != nil {
		return err
	}

	storeInitialized = true
	return nil
}

func InitClient() error {
	if clientInitialized {
		return nil
	}

	if err := InitStore(); err != nil {
		return err
	}

	Client = whatsmeow.NewClient(Store, loggers.CLLogger)

	// Authentication.
	if Client.Store.ID == nil {
		// No ID stored, this may be a new auth.
		qrChan, _ := Client.GetQRChannel(context.Background())
		if err := Client.Connect(); err != nil {
			return err
		}

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
		if err := Client.Connect(); err != nil {
			return err
		}
	}

	clientInitialized = true

	return nil
}
