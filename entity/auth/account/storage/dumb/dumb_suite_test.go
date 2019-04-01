package dumb_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDumb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Account Dumb Storage Suite")
}
