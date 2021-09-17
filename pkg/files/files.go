package files

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/jffin/multi-check-if-email-exists/pkg/checker"
)

func ReadInputFile(fileName string) []string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("reading input file %v", err)
	}
	return cleanupInputContent(content)
}

func WriteOutputFile(fileName string, results checker.Response) {
	data, _ := json.Marshal(results)
	if err := os.WriteFile(fileName, data, 0644); err != nil {
		log.Fatalf("writing to output file %v", err)
	}
}

func cleanupInputContent(content []byte) []string {
	var windowsSupportedString = strings.ReplaceAll(string(content), "\r\n", "\n")
	var withRemovedLastEmptyLine = strings.TrimRight(windowsSupportedString, "\n")
	return strings.Split(withRemovedLastEmptyLine, "\n")
}
