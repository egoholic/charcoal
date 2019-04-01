package mongodb_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMongodb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Account MongoDB Suite")
}
