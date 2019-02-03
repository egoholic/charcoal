package main_test

import (
	"fmt"
	"net/http"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var (
	err               error
	target            = "github.com/egoholic/charcoal/services/auth/target/http/main.go"
	buildPath         string
	runServiceCommand string
)

var _ = Describe("Auth HTTP servise", func() {
	BeforeSuite(func() {
		buildPath, err = gexec.Build(target)
		runServiceCommand := exec.Command(buildPath)
		_, err = gexec.Start(runServiceCommand, GinkgoWriter, GinkgoWriter)
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	Describe("GET /signup", func() {
		It("renders signup form successfully", func() {
			var response *http.Response
			response, err = http.Get("http://localhost:8080/test")
			if err != nil {
				fmt.Printf("\n\n\n\tERROR: %s\n\n", err.Error())
			}
			Expect(response).NotTo(BeNil())
			Expect(response.Status).To(Equal(""))
			Expect(response.Header).To(Equal(nil))
			Expect(response.Body).To(Equal(""))
		})
	})
})
