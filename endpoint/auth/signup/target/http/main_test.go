package main_test

import (
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/ghttp"
)

var (
	err                   error
	target                string
	buildPath             string
	runServiceCommand     string
	server                = ghttp.NewServer()
	session               *gexec.Session
	login1                = "email@example.com"
	password1             = "pwd12345678"
	passwordConfirmation1 = "pwd12345678"
)

type postPayload struct {
	login                string
	password             string
	passwordConfirmation string
}

func (pp *postPayload) Read(p []byte) (n int, err error) {
	destinationLen := len(p)
	data := []byte(pp.JSON())
	sourceLen := len(data)

	if len(data) == 0 {
		n = 0
		return
	}

	var i = 0
	for ; i < sourceLen; i++ {
		n = i + 1

		if i >= sourceLen || i >= destinationLen {
			return
		}

		p[i] = data[i]
	}

	return
}

func (pp *postPayload) JSON() string {
	result := &strings.Builder{}
	result.WriteRune('{')
	result.WriteString("\"login\":\"")
	result.WriteString(pp.login)
	result.WriteString("\",")

	result.WriteString("\"password\":\"")
	result.WriteString(pp.password)
	result.WriteString("\",")

	result.WriteString("\"passwordConfirmation\":\"")
	result.WriteString(pp.passwordConfirmation)
	result.WriteString("\"")

	result.WriteRune('}')
	return result.String()
}

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

		Expect(err).ShouldNot(HaveOccurred())

		//session.Wait(5 * time.Second)
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
		session.Interrupt()
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
			fmt.Printf("\n\n\n%s\n\n\n", logs)
			expectedLogs := "[ auth-1 ] Service initialization!"
			Expect(logs).To(Equal(expectedLogs))
		})
	})

	Describe("POST /signup", func() {
		Context("when correct creds", func() {
			It("creates new account, signs user in and responds with success", func() {
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("POST", "/signup", ""),
						ghttp.VerifyHeader(http.Header{
							"Content-Length": []string{"218"},
						}),
						ghttp.RespondWith(http.StatusOK, ""),
					),
				)

				payload := &postPayload{login1, password1, passwordConfirmation1}
				_, err := http.Post("http://localhost:8080/signup", "application/json", payload)
				Expect(err).NotTo(HaveOccurred())

				logs := string(session.Wait(500 * time.Millisecond).Out.Contents())
				expectedLogs := "Auth service started!\nlisten tcp :8080: bind: address already in use\n"
				fmt.Printf("\n\n\n%s\n\n\n", logs)

				Expect(logs).To(Equal(expectedLogs))
			})
		})
	})
})
