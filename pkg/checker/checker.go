package checker

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os/exec"
)

const (
	RequestUrl         string = "http://127.0.0.1:3000"
	RequestContentType string = "application/json"

	AppToRun   string = "app/check_if_email_exists"
	AppOptions string = "--http"
)

type Request struct {
	ToEmails []string `json:"to_emails"`
}

type MXResponse struct {
	AcceptsMail bool     `json:"accepts_mail"`
	Records     []string `json:"records"`
}

type SyntaxResponse struct {
	Address       string `json:"address"`
	Domain        string `json:"domain"`
	IsValidSyntax bool   `json:"is_valid_syntax"`
	Username      string `json:"username"`
}

type Response []struct {
	Input       string            `json:"input"`
	IsReachable string            `json:"is_reachable"`
	Misc        map[string]string `json:"misc"`
	Mx          MXResponse        `json:"mx"`
	Smtp        map[string]bool   `json:"smtp"`
	Syntax      SyntaxResponse    `json:"syntax"`
}

func Check(targetsArray []string) Response {
	command := startRustCheck()

	postBody, _ := json.Marshal(Request{ToEmails: targetsArray})
	requestBody := bytes.NewBuffer(postBody)

	response, err := http.Post(RequestUrl, RequestContentType, requestBody)

	if err != nil {
		log.Fatalf("An Error Occurred %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Failed to close %v", err)
		}
	}(response.Body)
	killRustCheck(command)

	return readResponse(response)
}

func startRustCheck() *exec.Cmd {
	cmd := exec.Command(AppToRun, AppOptions)
	if err := cmd.Start(); err != nil {
		log.Fatalf("check_if_email_exists fail to start with error: %v", err)
	}
	return cmd
}

func killRustCheck(command *exec.Cmd) {
	if err := command.Process.Kill(); err != nil {
		log.Fatalf("failed to kill process: %v", err)
	}
}

func readResponse(response *http.Response) Response {
	var b bytes.Buffer
	if _, err := io.Copy(&b, response.Body); err != nil {
		log.Fatalln("reading response body", err)
	}

	var responseData Response
	if err := json.Unmarshal([]byte(b.String()), &responseData); err != nil {
		log.Fatalf("Unsacsesful deserialization %v", err)
	}

	return responseData
}
