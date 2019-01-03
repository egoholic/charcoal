package charcoal_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCharcoal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Charcoal Suite")
}
