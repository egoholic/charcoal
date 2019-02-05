package main_test

import (
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var (
	err               error
	target            string
	buildPath         string
	runServiceCommand string
	session           *gexec.Session
)

var _ = Describe("Auth HTTP servise", func() {
	BeforeSuite(func() {
		target, err = filepath.Abs("./main.go")
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		buildPath, err = gexec.Build(target)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		runServiceCommand := exec.Command(buildPath)
		session, err = gexec.Start(runServiceCommand, GinkgoWriter, GinkgoWriter)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	Describe("GET /signup", func() {
		It("renders signup form successfully", func() {
			var response *http.Response
			response, err = http.Get("http://localhost:8080/signup")
			Expect(err).NotTo(HaveOccurred())
			Expect(response).NotTo(BeNil())
			Expect(response.Status).To(Equal("200 OK"))
			Expect(response.Header.Get("Content-Length")).To(Equal("3"))
			var body []byte
			response.Body.Read(body)
			Expect(string(body)).To(Equal("hey"))
			response.Body.Close()
			Expect(session.Wait(500 * time.Millisecond).Out.Contents()).To(Equal(""))
		})
	})
})
