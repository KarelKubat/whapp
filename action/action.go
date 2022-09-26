package action

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type Info struct {
	Name  string // name of an action
	NArgs int    // supported nr of args
	Usage string // short info
}

type Action interface {
	Description() *Info
	Run(args []string) error
}

var (
	registry []Action
	mu       sync.Mutex
)

func Register(a Action) {
	mu.Lock()
	defer mu.Unlock()
	registry = append(registry, a)
}

func Usage() string {
	desc := []string{}
	for _, a := range registry {
		desc = append(desc, a.Description().Usage)
	}
	sort.Strings(desc)
	return strings.Join(desc, "\n")
}

func Run(args []string) error {
	for _, a := range registry {
		if a.Description().Name == args[0] {
			if a.Description().NArgs != len(args)-1 {
				return fmt.Errorf("bad nr of args for %q, usage: %v", args[0], a.Description().Usage)
			}
			return a.Run(args[1:])
		}
	}
	return fmt.Errorf("no such action %q, choose one of:\n%v", args[0], Usage())
}
