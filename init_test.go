package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Suite")
}

var _ = Describe("Book", func() {

	It("Creates", func() {
		init := InitTask{
			ImportPath: "path",
		}
		Expect(init.ImportPath).To(Equal("path"))
	})
})
