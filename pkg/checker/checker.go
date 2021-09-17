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
	requestUrl         string = "http://127.0.0.1:3000"
	requestContentType string = "application/json"

	appToRun   string = "app/check_if_email_exists"
	appOptions string = "--http"
)

type request struct {
	ToEmails []string `json:"to_emails"`
}

type mxResponse struct {
	AcceptsMail bool     `json:"accepts_mail"`
	Records     []string `json:"records"`
}

type syntaxResponse struct {
	Address       string `json:"address"`
	Domain        string `json:"domain"`
	IsValidSyntax bool   `json:"is_valid_syntax"`
	Username      string `json:"username"`
}

type Response []struct {
	Input       string            `json:"input"`
	IsReachable string            `json:"is_reachable"`
	Misc        map[string]string `json:"misc"`
	Mx          mxResponse        `json:"mx"`
	Smtp        map[string]bool   `json:"smtp"`
	Syntax      syntaxResponse    `json:"syntax"`
}

func Check(targetsArray []string) Response {
	command := startRustCheck()

	postBody, _ := json.Marshal(request{ToEmails: targetsArray})
	requestBody := bytes.NewBuffer(postBody)

	response, err := http.Post(requestUrl, requestContentType, requestBody)

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
	cmd := exec.Command(appToRun, appOptions)
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
		log.Fatalf("Unsuccessful deserialization %v", err)
	}

	return responseData
}
