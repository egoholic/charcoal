package signin_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSignin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Signin Suite")
}
