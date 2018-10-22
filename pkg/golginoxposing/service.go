package golginoxposing

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
	"gitlab.com/TheDonDope/gocha/v3/pkg/errors"
	"gitlab.com/TheDonDope/gocha/v3/pkg/logging"
	"gitlab.com/TheDonDope/twitter-frequenter/pkg/config"
	storage "gitlab.com/TheDonDope/twitter-frequenter/pkg/storage/csv"
)

// Service describes methods
type Service interface {
	// Execute the service method, either concurrently or serially based on configration options.
	Execute()
	// Concurrently execute leverages Go routines and channels during runtime.
	Concurrently()
	// Serially execute processes the data in a serial manner.
	Serially()
}

type service struct{}

// NewService returns a new implementation instance for the Service interface
func NewService() Service {
	return &service{}
}

// Execute the service method, either concurrently or serially based on configration options.
func (s *service) Execute() {
	if config.Opts.Concurrent {
		logging.Printfln("Starting concurrent processing @ %v", time.Now().Format(time.RFC3339))
		s.Concurrently()
		logging.Printfln("Finished concurrent processing @ %v", time.Now().Format(time.RFC3339))
	} else {
		logging.Printfln("Starting non-concurrent processing @ %v", time.Now().Format(time.RFC3339))
		s.Serially()
		logging.Printfln("Finished non-concurrent processing @ %v", time.Now().Format(time.RFC3339))
	}
}

// Concurrently leverages Go routines and channels during runtime.
func (s *service) Concurrently() {

	if _, err := os.Stat(config.Opts.File); os.IsNotExist(err) {
		errors.Print(err, "abort: file '"+config.Opts.File+"' does not seem to exist")
	}

	csvFile, err := os.Open(config.Opts.File)
	if err != nil {
		errors.Print(err, "Unable to open file: "+config.Opts.File)
	}
	defer func() {
		if err := csvFile.Close(); err != nil {
			errors.Print(err, "Unable to close file: "+config.Opts.File)
			panic(err)
		}
	}()
	csvReader := csv.NewReader(csvFile)
	decoder, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		errors.Print(err, "Error creating decoder")
	}

	if "user" == config.Opts.Task {
		parser := storage.NewUserParser()
		// Nr. of workers = cpu core count - 1 for the main go routine.
		numWorkers := int(math.Max(1.0, float64(runtime.NumCPU()-1)))

		// see: https://github.com/jszwec/csvutil#unmarshal-and-metadata-

		users, decoded, done := make(chan []storage.User, numWorkers), make(chan storage.User, numWorkers), make(chan int)

		// Start the number of workers (parsers) determined by numWorkers.
		logging.Printfln("Starting %v workers...", numWorkers)
		for i := 0; i < numWorkers; i++ {
			go parser.ParseSnippet(i, users, decoded, done)
		}

		// Main file scanner go routine.
		go func() {
			defer func() {
				close(users)
			}()

			// header := decoder.Header()
			users := []storage.User{}
			for {
				user := storage.User{}
				if err := decoder.Decode(&user); err == io.EOF {
					break
				} else if err != nil {
					log.Fatal(err)
				}
				users = append(users, user)
				decoded <- user
			}
		}()

		var allUsers []storage.User
		// Wait for parsed snippets to come in, and ensure their uniqueness by using a map.
		waits := numWorkers
		for {
			select {
			case user := <-decoded:
				logging.Printfln("Append user: %v", user)
				allUsers = append(allUsers, user)
			case <-done:
				waits--
				if waits == 0 {
					logging.Printfln("Finished user waits @ %v", time.Now().Format(time.RFC3339))
					return
				}
			}
		}
	} else if "tweet" == config.Opts.Task {
		// parser := api.NewTweetParser()
	}

}

// Serially execute processes the data in a serial manner.
func (s *service) Serially() {
	if "user" == config.Opts.Task {
		parser := storage.NewUserParser()
		csvBytes, csvError := ioutil.ReadFile(config.Opts.File)
		errors.Print(csvError, "Error opening CSV file for path: "+config.Opts.File)
		parser.FromCSV(csvBytes)
	} else if "tweet" == config.Opts.Task {
		parser := storage.NewTweetParser()
		csvBytes, csvError := ioutil.ReadFile(config.Opts.File)
		errors.Print(csvError, "Error opening CSV file for path: "+config.Opts.File)
		parser.FromCSV(csvBytes)
	}
}
