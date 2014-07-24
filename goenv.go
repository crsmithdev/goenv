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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const template = `
# This file must be used with "source bin/activate" or ". bin/activate"

if [[ -n "${GOENV+1}" ]]; then
	deactivate
fi

export GOENV=__GOENV__
export GOENV_OLDPS1=$PS1
export GOENV_OLDGOPATH=$GOPATH
export GOENV_OLDPATH=$PATH

export GOPATH=$GOENV:$GOPATH
export PATH="$GOENV/bin:$PATH"
export PS1="($(basename $GOENV))$PS1"

deactivate() {
	export PS1=$GOENV_OLDPS1
	export GOPATH=$GOENV_OLDGOPATH
	export PATH=$GOENV_OLDPATH

	unset GOENV GOENV_OLDPS1 GOENV_OLDPATH GOENV_OLDGOPATH
	unset -f deactivate
}
`

// writeScript writes the modifed script template.
func writeScript(goenv, path string) error {

	script := strings.Replace(template, "__GOENV__", goenv, -1)
	err := ioutil.WriteFile(path, []byte(script), 777)

	return err
}

// createSubdirs creates the goenv directory structure.
func createSubdirs(path string, names []string) error {

	for _, name := range names {
		subDir := filepath.Join(path, name)
		if err := os.MkdirAll(subDir, os.ModeDir|0755); err != nil {
			return err
		}
	}

	return nil
}

func init() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: goenv DEST_DIR\n")
	}
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	dir, err := os.Getwd()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current directory: %s", err.Error())
		os.Exit(1)
	}

	path := filepath.Join(dir, flag.Arg(0))
	err = createSubdirs(path, []string{"src", "bin", "pkg"})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directories: %s", err.Error())
		os.Exit(1)
	}

	scriptPath := filepath.Join(path, "bin/activate")
	err = writeScript(path, scriptPath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing activate script: %s", err.Error())
		os.Exit(1)
	}
}
