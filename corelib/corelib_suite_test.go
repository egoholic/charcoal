package corelib_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCorelib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Corelib Suite")
}
