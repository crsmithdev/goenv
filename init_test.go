package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

func TestBooks(t *testing.T) {
	RunSpecs(t, "Init")
}

var _ = Describe("Init", func() {

	It("Creates an init task with defaults", func() {

		usr, _ := user.Current()
		wd, _ := os.Getwd()

		init, err := NewInitTask([]string{"test/path"})
		initTask := init.(*InitTask)

		assert.Nil(GinkgoT(), err)
		assert.Equal(GinkgoT(), initTask.ImportPath, "test/path")
		assert.Equal(GinkgoT(), initTask.ProjectPath, wd)
		assert.Equal(GinkgoT(), initTask.ProjectName, "goenv")
		assert.Equal(GinkgoT(), initTask.GoPath, usr.HomeDir+"/.goenv/goenv")
	})

	It("Creates an init task with arguments", func() {

		init, err := NewInitTask([]string{
			"-n",
			"name",
			"-g",
			"gopath",
			"-s",
			"activate",
			"-p",
			"path",
			"test/path",
		})

		initTask := init.(*InitTask)

		assert.Nil(GinkgoT(), err)
		assert.Equal(GinkgoT(), initTask.ImportPath, "test/path")
		assert.Equal(GinkgoT(), initTask.ProjectPath, "path")
		assert.Equal(GinkgoT(), initTask.ProjectName, "name")
		assert.Equal(GinkgoT(), initTask.GoPath, "gopath")
	})

	It("Returns an error if no import path is given", func() {

		_, err := NewInitTask([]string{"-n", "name"})

		assert.NotNil(GinkgoT(), err)
		assert.Contains(GinkgoT(), err.Error(), "import")
	})

	It("Generates an activate script with the specified options", func() {

		wd, _ := os.Getwd()
		script := filepath.Join(wd, "temp")

		task := &InitTask{
			GoPath:      "go",
			ImportPath:  "import/path",
			ProjectName: "name",
			ProjectPath: "path",
			ScriptPath:  script,
		}

		task.Run()

		bytes, _ := ioutil.ReadFile(script)
		text := string(bytes)

		assert.Contains(GinkgoT(), text, "export GOENV=name")
		assert.Contains(GinkgoT(), text, "export GOPATH=go")
		assert.Contains(GinkgoT(), text, "mkdir -p $(dirname $GOPATH/src/import/path")
		assert.Contains(GinkgoT(), text, "ln -s path $GOPATH/src/import/path")

		err := os.Remove(script)

		fmt.Println(err)
	})
})
