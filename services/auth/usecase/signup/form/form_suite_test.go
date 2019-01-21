package form_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestForm(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Form Suite")
}
