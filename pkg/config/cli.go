package config

// Opts are the program options, configurable by command line argument
var Opts struct {
	Concurrent bool `short:"c" long:"concurrent" description:"Run golginoxpose with enabled concurrency."`

	File string `short:"f" long:"file" description:"The path for the input CSV file to read."`

	OutputDirectory string `short:"o" long:"output-directory" description:"The output directory for the log file, results.json and dump.json (default: output)" default:"output"`

	Task string `short:"t" long:"task" description:"The task to execute." choice:"user" choice:"tweet"`
}

// GetOutName returns the custom file/directory name
func GetOutName() string {
	result := ""
	if Opts.File != "" {
		result = "file-" + Opts.Task
	}
	return result
}

// GetOutPath returns the complete path to the output directory
func GetOutPath() string {
	return Opts.OutputDirectory + "/" + GetOutName()
}
