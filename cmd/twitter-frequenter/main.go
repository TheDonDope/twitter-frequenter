package main

import (
	"os"

	"github.com/TheDonDope/twitter-frequenter/pkg/config"
	"github.com/TheDonDope/twitter-frequenter/pkg/golginoxposing"
)

func main() {
	// TODO: rewrite
	// cli.ParseArgs(&config.Opts, os.Args)
	os.MkdirAll(config.GetOutPath(), 0700)
	// TODO: rewrite
	// f := cli.NewLogFile(config.GetOutPath() + "/" + config.GetOutName() + ".log")
	// defer f.Close()

	// TODO: rewrite
	// start := time.Now()
	// logging.Printfln("Starting golginoxpose @ %v", time.Now().Format(time.RFC3339))
	golgi := golginoxposing.NewService()
	golgi.Execute()
	// TODO: rewrite
	// logging.Printfln("Program arguments: %+v", config.Opts)
	// logging.Printfln("Finishing golginoxpose @ %v", time.Now().Format(time.RFC3339))
	// logging.Printfln("Overall time spent: %v", time.Since(start))
}
