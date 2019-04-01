package value_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestValue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Value Suite")
}
