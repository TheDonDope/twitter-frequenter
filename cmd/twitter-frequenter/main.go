package main

import (
	"os"
	"time"

	cli "gitlab.com/TheDonDope/gocha/v3/pkg/config"
	"gitlab.com/TheDonDope/gocha/v3/pkg/logging"
	"gitlab.com/TheDonDope/twitter-frequenter/pkg/config"
	"gitlab.com/TheDonDope/twitter-frequenter/pkg/golginoxposing"
)

func main() {
	cli.ParseArgs(&config.Opts, os.Args)
	os.MkdirAll(config.GetOutPath(), 0700)
	f := cli.NewLogFile(config.GetOutPath() + "/" + config.GetOutName() + ".log")
	defer f.Close()

	start := time.Now()
	logging.Printfln("Starting golginoxpose @ %v", time.Now().Format(time.RFC3339))
	golgi := golginoxposing.NewService()
	golgi.Execute()
	logging.Printfln("Program arguments: %+v", config.Opts)
	logging.Printfln("Finishing golginoxpose @ %v", time.Now().Format(time.RFC3339))
	logging.Printfln("Overall time spent: %v", time.Since(start))
}
