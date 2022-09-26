package flags

import "flag"

var (
	LogfileFlag   = flag.String("logfile", "/tmp/whapp.log", "logfile to write")
	VerboseFlag   = flag.Bool("verbose", false, "when true, debug messages are logged")
	AppendFlag    = flag.Bool("append", true, "when true (default), the logfile is appended")
	DBFlag        = flag.String("db", "store.db", "sqlite3 backend")
	QuitAfterFlag = flag.Duration("quit-after", 0, "stop after the given duration (default: poll forever")
)
