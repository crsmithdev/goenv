/*
goenv provides isolated, virtual GOPATH environments for Go projects.

Usage:

    goenv [command] [arguments]

Commands:

    init     initialize a goenv

Use "goenv help [command]" for command-specific information.

goenv: unrecognized command init
*/
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"text/template"
)

var usageTemplate = `
Goenv provides isolated, virtual GOPATH environments for Go projects.

Usage:

    goenv [command] [arguments]

Commands:
{{ range . }}
    {{ .Name | printf "%-8s" }} {{ .Short }}{{end}}

Use "goenv help [command]" for command-specific information.
`

// Command is a command-line action.
type Command struct {
	Name    string
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
var usageText string

func init() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: goenv DEST_DIR\n")
	}

	commands := make(map[string]*Command)

	cmds := []*Command{
		&initCommand,
	}

	for _, cmd := range cmds {
		commands[cmd.Name] = cmd
	}

	tmpl := template.New("usage")
	tmpl, err := tmpl.Parse(usageTemplate)

	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, cmds)

	if err != nil {
		panic(err)
	}

	usageText = buf.String()
	fmt.Println(usageText)

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
