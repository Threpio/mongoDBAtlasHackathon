package main

import (
	"errors"
	"fmt"
	"os"
)

var versionNumber = "0.0.1"

var cliHelp = `Command-line interface and server.
Usage core [command]
Available Commands:
	start
	version
Use "core [command] --help" for more information about a command.
`

func Run(args []string) error {
	switch args[0] {
	case "start":
		StartServer()
		// TODO: Start Function
		return errors.New("not implemented")
	case "version":
		return version()
	default:
		return help()
	}
}

func version() error {
	fmt.Printf("Version: %s", versionNumber)
	return nil
}

func help() error {
	return errors.New(cliHelp)
}

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "help")
	}

	if err := Run(os.Args[1:]); err != nil {
		fmt.Println(err)
	}
}
