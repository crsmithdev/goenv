// Test3.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var template = `
# This file must be used with "source bin/activate" or ". bin/activate"

GOENV_NAME="__GOENV_NAME__"
GOENV_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

if [[ "$GOENV_ACTIVE" == "1" ]]; then
	deactivate
fi

export GOENV_ACTIVE=1
export GOENV_PS1=$PS1
export GOENV_GOPATH=$GOPATH
export GOENV_PATH=$PATH

export GOPATH="$GOENV_DIR/$GOENV_NAME:$GOPATH"
export PATH="$GOENV_DIR/bin:$PATH"
export PS1="(${GOENV_NAME})${PS1}"

deactivate() {
	export PS1=$GOENV_PS1
	export GOPATH=$GOENV_GOPATH
	export PATH=$GOENV_PATH

	unset GOENV_ACTIVE GOENV_PS1 GOENV_PATH GOENV_GOPATH
	unset -f deactivate
}
`

func writeScript(path, name string) error {

	script := strings.Replace(template, "__GOENV_NAME__", name, -1)
	err := ioutil.WriteFile(path, []byte(script), 777)

	return err
}

func createSubdirs(path string, names []string) error {

	for _, name := range names {
		subDir := filepath.Join(path, name)
		if err := os.MkdirAll(subDir, os.ModeDir|0755); err != nil {
			return err
		}
	}

	return nil
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: goenv DEST_DIR\n")
	}

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	path := flag.Arg(0)
	err := createSubdirs(path, []string{"src", "bin", "pkg"})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory structure: %s", err)
		os.Exit(1)
	}

	scriptPath := filepath.Join(path, "bin/activate")
	err = writeScript(scriptPath, filepath.Base(path))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing activate script: %s", err)
		os.Exit(1)
	}
}
