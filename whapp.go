package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/KarelKubat/flagnames"
	"github.com/KarelKubat/whapp/action"
	"github.com/KarelKubat/whapp/loggers"
	"github.com/KarelKubat/whapp/wm"

	// Actions that register themselves.
	_ "github.com/KarelKubat/whapp/action/poll"
)

const (
	usageInfo = `
Usage: whapp [FLAGS] ACTION [OPTIONAL-ARGS]

The actions (and optional arguments are):

%v

The accepted flags are listed below. They can be abbreviated (e.g., -v for -verbose).

`
)

func main() {
	// Commandline
	parseCommandLine()

	// Logging
	checkErr(loggers.Init())

	// Instantiate storage, that's always needed. Client initialization is left to the actions
	// that need it.
	checkErr(wm.InitStore())

	// Run the action of arg[1]
	action.Run(flag.Args())

}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: %v\n", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, usageInfo, action.Usage())
	flag.PrintDefaults()
	fmt.Fprintln(os.Stdout)
}

func parseCommandLine() {
	flagnames.Patch()
	flag.Usage = usage
	flag.Parse()

	// There must be at least one action arg.
	if flag.NArg() == 0 {
		usage()
	}
}
