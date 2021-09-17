package args

import "flag"

const (
	defaultTargetArg string = "target"
	defaultInputArg  string = "input"
	defaultOutputArg string = "output"
	TARGET           string = ""
	inputFileName    string = "input.txt"
	resultFileName   string = "results.json"
)

func InitArgs() (target, inputFile, outputFile *string) {
	target = flag.String(defaultTargetArg, TARGET, "Single target")
	inputFile = flag.String(defaultInputArg, inputFileName, "Input file with targets")
	outputFile = flag.String(defaultOutputArg, resultFileName, "Output file write results to")
	flag.Parse()

	return
}
