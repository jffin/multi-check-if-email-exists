package args

import "flag"

const (
	DefaultTargetArg   string = "target"
	DefaultInputArg  string = "input"
	DefaultOutputArg string = "output"
	TARGET           string = ""
	InputFileName  string = "input.txt"
	ResultFileName string = "results.json"
)

func InitArgs() (target, inputFile, outputFile *string) {
	target = flag.String(DefaultTargetArg, TARGET, "Single target")
	inputFile = flag.String(DefaultInputArg, InputFileName, "Input file with targets")
	outputFile = flag.String(DefaultOutputArg, ResultFileName, "Output file write results to")
	flag.Parse()

	return
}
