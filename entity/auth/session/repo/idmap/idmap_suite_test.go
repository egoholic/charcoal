package idmap_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIdmap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Session Idmap Suite")
}
