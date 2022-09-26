package loggers

import (
	"github.com/KarelKubat/whapp/flags"
	"github.com/KarelKubat/whatsmeow/logger"

	waLog "go.mau.fi/whatsmeow/util/log"
)

var (
	DBLogger waLog.Logger
	CLLogger waLog.Logger
)

func Init() error {
	baseLogger, err := logger.New(logger.Opts{
		Filename: *flags.LogfileFlag,
		Verbose:  *flags.VerboseFlag,
		Append:   *flags.AppendFlag,
	})
	if err != nil {
		return err
	}
	DBLogger = baseLogger.Sub("Database")
	CLLogger = baseLogger.Sub("Client")
	return nil
}
