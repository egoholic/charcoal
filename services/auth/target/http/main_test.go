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
	"github.com/onsi/gomega/ghttp"
)

var (
	err               error
	target            string
	buildPath         string
	runServiceCommand string
	server            = ghttp.NewServer()
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
		session.Wait(10 * time.Second)
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	Describe("GET /signup", func() {
		It("renders signup form successfully", func() {
			expectedBody := `<div id="sign-up">
          <h1>Sign Up</h1>

          <form action="/signup/" method="POST">
              <input type="text" name="login">
              <input type="text" name="password">
              <input type="submit">
          </form>
				</div>`

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/signup", ""),
					ghttp.VerifyHeader(http.Header{
						"Content-Length": []string{"218"},
					}),
					ghttp.RespondWith(http.StatusOK, expectedBody),
				),
			)

			_, err := http.Get("http://localhost:8080/signup")
			Expect(err).NotTo(HaveOccurred())
			logs := string(session.Wait(500 * time.Millisecond).Out.Contents())
			expectedLogs := "Auth service started!\nlisten tcp :8080: bind: address already in use\n"
			Expect(logs).To(Equal(expectedLogs))
		})
	})
})
