package signup_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSignup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Signup Suite")
}
