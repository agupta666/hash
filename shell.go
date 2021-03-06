package main

import (
	"fmt"
	"os"

	"github.com/agupta666/elf/actions"
	"github.com/agupta666/elf/commands"
	"github.com/agupta666/elf/router"
	shellwords "github.com/mattn/go-shellwords"
	readline "gopkg.in/readline.v1"
)

func processCmd(line string) {
	args, err := shellwords.Parse(line)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: syntax error")
		return
	}

	if len(args) > 0 {
		handler := commands.LookupHandler(args[0])
		handler(args[1:])
	}
}

func startShell() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "elf> ",
		HistoryFile:  ".elf.hist",
		AutoComplete: completer,
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		processCmd(line)
	}
}

var completer = readline.NewPrefixCompleter(
	readline.PcItem("route",
		readline.PcItem("/path", readline.PcItemDynamic(actions.ActionList)),
	),
	readline.PcItem("lsrt"),
	readline.PcItem("delrt", readline.PcItemDynamic(router.RouteNames)),
	readline.PcItem("kvset"),
	readline.PcItem("lskv"),
	readline.PcItem("help", readline.PcItemDynamic(commands.CommandList)),
	readline.PcItem("quit"),
)
