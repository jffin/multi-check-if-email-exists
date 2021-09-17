package args

import "flag"

const (
	DEFAULT_TARGET_ARG string = "target"
	DEFAULT_INPUT_ARG  string = "input"
	DEFAULT_OUTPUT_ARG string = "output"
	TARGET             string = ""
	INPUT_FILE_NAME    string = "input.txt"
	RESULT_FILE_NAME   string = "results.json"
)

func InitArgs() (target, inputFile, outputFile *string) {
	target = flag.String(DEFAULT_TARGET_ARG, TARGET, "Single target")
	inputFile = flag.String(DEFAULT_INPUT_ARG, INPUT_FILE_NAME, "Input file with targets")
	outputFile = flag.String(DEFAULT_OUTPUT_ARG, RESULT_FILE_NAME, "Output file write results to")
	flag.Parse()

	return
}
