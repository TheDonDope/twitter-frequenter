package api

import (
	"gitlab.com/TheDonDope/twitter-frequenter/pkg/types"
)

// UserParser defines methods to read twitter user data from a source CSV
type UserParser interface {
	// FromCSV returns an array of TwitterUsers read from a CSV file
	FromCSV([]byte) ([]types.User, error)

	// ParseSnippet parses ...
	ParseSnippet(int, <-chan []types.User, chan<- types.User, chan<- int)
}
