package api

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"runtime"
	"time"

	"github.com/jszwec/csvutil"
	"gitlab.com/TheDonDope/twitter-frequenter/pkg/types"
	"gitlab.com/TheDonDope/twitter-frequenter/pkg/util/configs"
	"gitlab.com/TheDonDope/twitter-frequenter/pkg/util/errors"
	"gitlab.com/TheDonDope/twitter-frequenter/pkg/util/logs"
)

// GolginoxposerService implements the Golginoxposer interface
type GolginoxposerService struct{}

// NewGolginoxposer returns a new implementation instance for the Golginoxposer interface
func NewGolginoxposer() Golginoxposer {
	return &GolginoxposerService{}
}

// ConcurrentExecute leverages Go routines and channels during runtime.
func (p GolginoxposerService) ConcurrentExecute() {

	if _, err := os.Stat(configs.Opts.File); os.IsNotExist(err) {
		errors.HandleError(err, "abort: file '"+configs.Opts.File+"' does not seem to exist")
	}

	csvFile, err := os.Open(configs.Opts.File)
	if err != nil {
		errors.HandleError(err, "Unable to open file: "+configs.Opts.File)
	}
	defer func() {
		if err := csvFile.Close(); err != nil {
			errors.HandleError(err, "Unable to close file: "+configs.Opts.File)
			panic(err)
		}
	}()
	csvReader := csv.NewReader(csvFile)
	decoder, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		errors.HandleError(err, "Error creating decoder")
	}

	if "user" == configs.Opts.Task {
		parser := NewUserParser()
		// Nr. of workers = cpu core count - 1 for the main go routine.
		numWorkers := int(math.Max(1.0, float64(runtime.NumCPU()-1)))

		// see: https://github.com/jszwec/csvutil#unmarshal-and-metadata-

		users, decoded, done := make(chan []types.User, numWorkers), make(chan types.User, numWorkers), make(chan int)

		// Start the number of workers (parsers) determined by numWorkers.
		logs.Printfln("Starting %v workers...", numWorkers)
		for i := 0; i < numWorkers; i++ {
			go parser.ParseSnippet(i, users, decoded, done)
		}

		// Main file scanner go routine.
		go func() {
			defer func() {
				close(users)
			}()

			// header := decoder.Header()
			users := []types.User{}
			for {
				user := types.User{}
				if err := decoder.Decode(&user); err == io.EOF {
					break
				} else if err != nil {
					log.Fatal(err)
				}
				users = append(users, user)
				decoded <- user
			}
		}()

		var allUsers []types.User
		// Wait for parsed snippets to come in, and ensure their uniqueness by using a map.
		waits := numWorkers
		for {
			select {
			case user := <-decoded:
				logs.Printfln("Append user: %v", user)
				allUsers = append(allUsers, user)
			case <-done:
				waits--
				if waits == 0 {
					logs.Printfln("Finished user waits @ %v", time.Now().Format(time.RFC3339))
					return
				}
			}
		}
	} else if "tweet" == configs.Opts.Task {
		// parser := api.NewTweetParser()
	}

}

// SerialExecute processes the data in a serial manner.
func (p GolginoxposerService) SerialExecute() {
	if "user" == configs.Opts.Task {
		parser := NewUserParser()
		csvBytes, csvError := ioutil.ReadFile(configs.Opts.File)
		errors.HandleError(csvError, "Error opening CSV file for path: "+configs.Opts.File)
		parser.FromCSV(csvBytes)
	} else if "tweet" == configs.Opts.Task {
		parser := NewTweetParser()
		csvBytes, csvError := ioutil.ReadFile(configs.Opts.File)
		errors.HandleError(csvError, "Error opening CSV file for path: "+configs.Opts.File)
		parser.FromCSV(csvBytes)
	}
}
