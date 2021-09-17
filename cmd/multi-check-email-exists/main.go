package main

import (
	"github.com/jffin/multi-check-if-email-exists/pkg/args"
	"github.com/jffin/multi-check-if-email-exists/pkg/checker"
	"github.com/jffin/multi-check-if-email-exists/pkg/files"
)

const (
	TargetCheckIndex int = 0
)

func main() {
	target, inputFile, outputFile := args.InitArgs()

	var targets []string
	targets = append(targets, *target)

	if targets[TargetCheckIndex] == args.TARGET {
		targets = files.ReadInputFile(*inputFile)
	}

	results := checker.Check(targets)
	files.WriteOutputFile(*outputFile, results)
}
