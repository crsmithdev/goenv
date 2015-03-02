/*
Package goenv provides a virtualenv-like GOPATH and PATH isolation for go.

Usage is as follows:

	Create a goenv:
		goenv [NAME]

	Activate the goenv:
		. [NAME]/bin/activate

	Deactivate the goenv:
		deactivate

Example:
	goenv local
	. local/bin/activate
	deactivate
*/
package main

import (
	"flag"
	"fmt"
	"os"
)

// Command is a command-line action.
type Command struct {
	Usage   string
	Short   string
	Long    string
	GetTask func([]string) (Task, error)
}

// Task is an action.
type Task interface {
	Run() error
}

// a map of command names -> commands.
var commands map[string]*Command

func init() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: goenv DEST_DIR\n")
	}

	commands = map[string]*Command{
		"init": &initCommand,
	}
}

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	cmd, found := commands[args[0]]

	if !found {
		fmt.Fprintf(os.Stderr, "goenv: unrecognized command %s\n", args[0])
		os.Exit(1)
	}

	task, err := cmd.GetTask(args[1:])

	if err != nil {
		fmt.Fprintf(os.Stderr, "goenv: failed to parse command %s\n", cmd)
		os.Exit(1)
	}

	err = task.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "goenv: error running command: %s\n", err)
		os.Exit(1)
	}
}
