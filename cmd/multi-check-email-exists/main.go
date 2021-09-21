package main

import (
	"github.com/jffin/multi-check-if-email-exists/pkg/args"
	"github.com/jffin/multi-check-if-email-exists/pkg/checker"
	"github.com/jffin/multi-check-if-email-exists/pkg/files"
)

const (
	targetCheckIndex int = 0
)

func main() {
	target, inputFile, outputFile := args.InitArgs()
	targets := setupTargets(*target, *inputFile)
	results := checker.Check(targets)
	files.WriteOutputFile(*outputFile, results)
}

func setupTargets(target string, inputFile string) (targets []string) {
	targets = append(targets, target)
	if targets[targetCheckIndex] == args.DefaultTarget {
		targets = files.ReadInputFile(inputFile)
	}

	return
}
