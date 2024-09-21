package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/Francis-Tao-jinjin/go-bitesize/cmd-app/sub-cmd-arch/cmd"
)

var errInvalidSubCommand = errors.New("Invalid sub-command specified")

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: mync [http|grpc] -h\n")
	cmd.HandleHttp(w, []string{"-h"})
	cmd.HandleGrpc(w, []string{"-h"})
}

func handleCommand(w io.Writer, args []string) error {
	var err error
	if len(args) < 1 {
		err = errInvalidSubCommand
	} else {
		switch args[0] {
		case "http":
			return cmd.HandleHttp(w, args[1:])
		case "grpc":
			return cmd.HandleGrpc(w, args[1:])
		case "-h":
			printUsage(w)
		case "-help":
			printUsage(w)
		default:
			return errInvalidSubCommand
		}
	}
	if errors.Is(err, cmd.ErrNoServerSpecified) || errors.Is(err, errInvalidSubCommand) {
		fmt.Fprintln(w, err)
		printUsage(w)
	}
	return err
}

func main() {
	err := handleCommand(os.Stdout, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
