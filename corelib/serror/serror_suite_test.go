package serror_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSerror(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Serror Suite")
}
