package commands

import (
	"fmt"
	"os"
)

// CmdHandler is a type for all command handler functions
type CmdHandler func(args []string)

func notFoundHandler(args []string) {
	fmt.Fprintln(os.Stderr, "ERROR:", "command not found")
}

var commandsMap = map[string]CmdHandler{
	"route": routeCmd,
	"lsrt":  lsroutesCmd,
	"delrt": deleteRouteCmd,
	"kvset": kvsetCmd,
	"lskv":  lskvCmd,
	"help":  helpCmd,
	"quit":  exitCmd,
}

// LookupHandler looks up a command handler by command name
func LookupHandler(cmd string) CmdHandler {
	handler, ok := commandsMap[cmd]

	if ok {
		return handler
	}
	return notFoundHandler
}

// CommandList returns a list of commands
func CommandList(arg string) []string {
	cmds := make([]string, 0)
	for cmd := range commandsMap {
		cmds = append(cmds, cmd)
	}
	return cmds
}
