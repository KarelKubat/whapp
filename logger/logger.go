package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	waLog "go.mau.fi/whatsmeow/util/log"
)

const (
	timeFormat = "15:04:05.000"
)

// Global vars for all loggers.
var (
	writer   io.Writer  // singleton to send output from all logger instances
	mu       sync.Mutex // to synchronize writing
	opened   bool       // only 1 instance supported
	filename string     // logfile
	openbits int        // os.OpenFile bitmask
	reopen   bool       // atomic writes
)

// Opts allows the caller to configure a logger.
type Opts struct {
	Module   string // logged module name
	Filename string // output filename
	Verbose  bool   // when true, debug messages are sent
	Append   bool   // when true, the logfile is appended, else it is overwritten
	Reopen   bool   // when true, the logfile is reopened when it disappears (useful for log rotation)
}

type logger struct {
	module  string
	verbose bool
}

func New(o Opts) (*logger, error) {
	if opened {
		if o.Filename != filename {
			return nil, fmt.Errorf("logger.New cannot open a second log %q (%q is already open)", o.Filename, filename)
		}
	} else {
		openbits = os.O_CREATE | os.O_WRONLY
		if o.Append {
			openbits |= os.O_APPEND
		}
		opened = true
		filename = o.Filename
		reopen = o.Reopen
		var err error
		writer, err = os.OpenFile(o.Filename, openbits, 0644)
		if err != nil {
			return nil, err
		}
	}
	return &logger{
		module:  o.Module,
		verbose: o.Verbose,
	}, nil
}

func (l *logger) Errorf(msg string, args ...interface{}) {
	output("ERROR", l.module, true, fmt.Sprintf(msg, args...))
}

func (l *logger) Warnf(msg string, args ...interface{}) {
	output("WARN", l.module, true, fmt.Sprintf(msg, args...))
}

func (l *logger) Infof(msg string, args ...interface{}) {
	output("INFO", l.module, true, fmt.Sprintf(msg, args...))
}

func (l *logger) Debugf(msg string, args ...interface{}) {
	output("DEBUG", l.module, l.verbose, fmt.Sprintf(msg, args...))
}

func (l *logger) Sub(module string) waLog.Logger {
	var newModule string
	switch {
	case l.module == "" && module != "":
		newModule = module
	case l.module != "" && module == "":
		newModule = l.module
	case l.module != "" && module != "":
		newModule = fmt.Sprintf("%s/%s", l.module, module)
	}

	return &logger{
		module:  newModule,
		verbose: l.verbose,
	}
}

func output(level, module string, send bool, msg string) {
	if !send {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	_, err := os.Stat(filename)
	if err != nil {
		writer, err = os.OpenFile(filename, openbits, 0644)
		if err != nil {
			panic(err) // There is no where to escalate the error, best we can do is panic.
		}
	}
	writer.Write([]byte(fmt.Sprintf("%s [%s %s] %s\n", time.Now().Format(timeFormat), module, level, msg)))
}
