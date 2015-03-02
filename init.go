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

export GOENV={{ .ProjectName }}
export GOENV_OLDPS1=$PS1
export GOENV_OLDGOPATH=$GOPATH
export GOENV_OLDPATH=$PATH

export GOPATH={{ .GoenvDir }}
export PATH="$GOPATH/bin:$PATH"
export PS1="($(basename $GOPATH))$PS1"

deactivate() {
	export PS1=$GOENV_OLDPS1
	export GOPATH=$GOENV_OLDGOPATH
	export PATH=$GOENV_OLDPATH

	unset GOENV GOENV_OLDPS1 GOENV_OLDPATH GOENV_OLDGOPATH
	unset -f deactivate
}
`

var initCommand = Command{
	Name:    "init",
	Short:   "initialize a goenv",
	Usage:   "init [-g][-s][-w] [project name] [import path]",
	Long:    "TODO",
	GetTask: NewInitTask,
}

// InitTask initializes a goenv.
type InitTask struct {
	ProjectName string // the name of the project, e.g. "goenv".
	ImportPath  string // the import path of the project, e.g. "github.com/crsmithdev/goenv"
	GoenvDir    string // the goenv dir to create, default "~/.goenv/[ProjectName]"
	ScriptDir   string // the directory in which to write the activate script, default "./goenv"
	WorkingDir  string // the current working directory.
}

// NewInitTask returns a new InitTask created with the specified command-line args.
func NewInitTask(args []string) (Task, error) {

	flags := flag.NewFlagSet("init", flag.ExitOnError)

	goenvDir := flags.String("workspace", "", "the workspace directory")
	scriptDir := flags.String("script", "./goenv/activate", "the activation script path")

	flags.Parse(args)
	args = flags.Args()

	if len(args) < 1 {
		return nil, errors.New("no project name specified")
	}

	if len(args) < 2 {
		return nil, errors.New("no project path specified")
	}

	workingDir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	return &InitTask{
		ProjectName: args[0],
		ImportPath:  args[1],
		GoenvDir:    *goenvDir,
		ScriptDir:   *scriptDir,
		WorkingDir:  workingDir,
	}, nil
}

// Run exeuctes the InitTask
func (task *InitTask) Run() error {

	fmt.Println("goenv: initializing...")

	if task.GoenvDir == "" {
		usr, err := user.Current()

		if err != nil {
			return err
		}

		task.GoenvDir = filepath.Join(usr.HomeDir, ".goenv/", task.ProjectName)
	}

	if err := task.makePaths(); err != nil {
		return err
	}

	if err := task.writeScript(); err != nil {
		return err
	}

	fmt.Println("goenv: done")

	return nil
}

// makePaths creates directory paths for a goenv.
func (task *InitTask) makePaths() error {

	srcPath := filepath.Join(task.GoenvDir, "src", task.ImportPath)

	var dirs = []string{
		filepath.Dir(task.ScriptDir),
		filepath.Join(task.GoenvDir, "bin"),
		filepath.Join(task.GoenvDir, "pkg"),
		srcPath,
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, os.ModeDir|0755)

		fmt.Printf("goenv: created directory %s\n", dir)

		if err != nil {
			return err
		}
	}

	err := os.Remove(srcPath)
	err = os.Symlink(task.WorkingDir, srcPath)

	fmt.Printf("goenv: symlinked %s -> %s\n", task.WorkingDir, srcPath)

	return err
}

// writeScript writes the goenv activate script.
func (task *InitTask) writeScript() error {

	scriptTemplate := template.New("test")
	scriptTemplate, err := scriptTemplate.Parse(script)

	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = scriptTemplate.Execute(&buf, task)

	if err != nil {
		return err
	}

	_ = os.Remove(task.ScriptDir)
	err = ioutil.WriteFile(task.ScriptDir, []byte(buf.String()), 777)

	fmt.Printf("goenv: wrote activation script at %s\n", task.ScriptDir)

	return err
}
