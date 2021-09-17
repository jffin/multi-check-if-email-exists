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
	REQUEST_URL          string = "http://127.0.0.1:3000"
	REQUEST_CONTENT_TYPE string = "application/json"

	APP_TO_RUN  string = "app/check_if_email_exists"
	APP_OPTIONS string = "--http"
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

	response, err := http.Post(REQUEST_URL, REQUEST_CONTENT_TYPE, requestBody)

	if err != nil {
		log.Fatalf("An Error Occurred %v", err)
	}
	defer response.Body.Close()
	killRustCheck(command)

	return readResponse(response)
}

func startRustCheck() *exec.Command {
	cmd := exec.Command(APP_TO_RUN, APP_OPTIONS)
	if err := cmd.Start(); err != nil {
		log.Fatalf("check_if_email_exists fail to start with error: %v", err)
	}
	go func() {
		err := cmd.Wait()
		log.Printf("Command finished with error: %v", err)
	}()
	return &cmd
}

func killRustCheck(command *exec.Command) {
	if err := *command.Process.Kill(); err != nil {
		log.Fatalf("failed to kill process: ", err)
	}
}

func readResponse(response *http.Response) Response {
	var b bytes.Buffer
	if _, err := io.Copy(&b, response.Body); err != nil {
		log.Fatalln("reading response body", err)
	}

	var responseData Response
	json.Unmarshal([]byte(b.String()), &responseData)

	return responseData
}
