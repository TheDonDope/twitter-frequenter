package api

import "gitlab.com/TheDonDope/twitter-frequenter/pkg/types"

// TweetParser defines methods to read tweet data from a source CSV
type TweetParser interface {
	// FromCSV returns an array of Tweets read from a CSV file
	FromCSV([]byte) ([]types.Tweet, error)

	// ParseSnippet parses ...
	ParseSnippet(int, <-chan []types.Tweet, chan<- types.Tweet, chan<- int)
}
