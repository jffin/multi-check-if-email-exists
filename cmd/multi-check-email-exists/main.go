package main

import (
	"fmt"

	"github.com/jffin/multi-check-if-email-exists/pkg/args"
	"github.com/jffin/multi-check-if-email-exists/pkg/checker"
	"github.com/jffin/multi-check-if-email-exists/pkg/files"
)

const (
	TARGET_CHECK_INDEX int = 0
)

func main() {
	target, inputFile, outputFile := args.InitArgs()

	var targets []string
	targets = append(targets, *target)

	if targets[TARGET_CHECK_INDEX] == args.TARGET {
		targets = files.ReadInputFile(*inputFile)
	}

	results := checker.Check(targets)
	fmt.Println(fmt.Println(results))

	files.WriteOutputFile(*outputFile, results)
}
