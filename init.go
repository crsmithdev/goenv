package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"text/template"
)

const script = `
# This file must be used with "source activate" or ". activate"

if [[ -n "${GOENV+1}" ]]; then
	deactivate
fi

export GOENV={{.ProjectName}}
export GOENV_OLDPS1=$PS1
export GOENV_OLDGOPATH=$GOPATH
export GOENV_OLDPATH=$PATH

export GOPATH={{.GoPath}}
export PATH="$GOPATH/bin:$PATH"
export PS1="($(basename $GOPATH))$PS1"

mkdir -p $(dirname $GOPATH/src/{{.ImportPath}})
rm -f $GOPATH/src/{{.ImportPath}}
ln -s {{.ProjectPath}} $GOPATH/src/{{.ImportPath}}

deactivate() {
	export PS1=$GOENV_OLDPS1
	export GOPATH=$GOENV_OLDGOPATH
	export PATH=$GOENV_OLDPATH

	unset GOENV GOENV_OLDPS1 GOENV_OLDPATH GOENV_OLDGOPATH
	unset -f deactivate
}
`

var initCommand = Command{
	Name:  "init",
	Short: "initialize a goenv",
	Usage: "init [-g][-s][-p][-n] [import path]",
	Long: `
Init initializes a goenv and creates an initialization script that
activates it.  This script creates, if needed, a GOPATH directory
structure, symlinks the project into that structure at the specified
input path, and then alters the current session's GOPATH environment
variable to point to it.

The goenv can be deactivated with 'deactivate'.

Init supports the following options:

    -n
         the name of the environment, defaulting to the name
         of the current working directory.

    -g
         the GOPATH to create, defaulting to ~/.goenv/<name>

    -p
         the project path, defaulting to the current working
         directory.

    -s
         the full path to the initialization script, defaulting
         to ./goenv/activate.
`,
	GetTask: NewInitTask,
}

// InitTask initializes a goenv.
type InitTask struct {
	GoPath      string // the GOPATH to create, default "~/.goenv/<project name>"
	ImportPath  string // the import path of the project, e.g. "github.com/crsmithdev/goenv"
	ProjectName string // the name of the project, e.g. "goenv".
	ProjectPath string // the path to the project, default "./"
	ScriptPath  string // the path and filename of the activation script to write, default "./goenv/activate"
}

// NewInitTask returns a new InitTask created with the specified command-line args.
func NewInitTask(args []string) (Task, error) {

	flags := flag.NewFlagSet("init", flag.ExitOnError)

	goPath := flags.String("g", "", "the GOPATH to create")
	projectName := flags.String("n", "", "the project name")
	scriptPath := flags.String("s", "./goenv/activate", "the full path to the initialization script")
	projectPath := flags.String("p", "", "the project path")

	flags.Parse(args)
	args = flags.Args()

	if len(args) < 1 {
		return nil, errors.New("no import path specified")
	}

	workingDir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	task := InitTask{
		ImportPath:  args[0],
		GoPath:      *goPath,
		ProjectName: *projectName,
		ProjectPath: *projectPath,
		ScriptPath:  *scriptPath,
	}

	if task.ProjectName == "" {
		task.ProjectName = filepath.Base(workingDir)
	}

	if task.ProjectPath == "" {
		task.ProjectPath = workingDir
	}

	if task.GoPath == "" {
		usr, err := user.Current()

		if err != nil {
			return nil, err
		}

		task.GoPath = filepath.Join(usr.HomeDir, ".goenv/", task.ProjectName)
	}

	return &task, nil
}

// Run exeuctes the InitTask
func (task *InitTask) Run() error {

	fmt.Println("goenv: initializing...")

	if err := task.writeScript(); err != nil {
		return err
	}

	fmt.Println("goenv: done")

	return nil
}

// writeScript writes the goenv activate script.
func (task *InitTask) writeScript() error {

	err := os.MkdirAll(filepath.Dir(task.ScriptPath), os.ModeDir|0755)

	if err != nil {
		return err
	}

	scriptTemplate := template.New("test")
	scriptTemplate, err = scriptTemplate.Parse(script)

	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = scriptTemplate.Execute(&buf, task)

	if err != nil {
		return err
	}

	_ = os.Remove(task.ScriptPath)
	err = ioutil.WriteFile(task.ScriptPath, []byte(buf.String()), 777)

	fmt.Printf("goenv: wrote activation script at %s\n", task.ScriptPath)

	return err
}
