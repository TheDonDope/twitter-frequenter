package api

import (
	"fmt"

	"github.com/jszwec/csvutil"
	"gitlab.com/TheDonDope/twitter-frequenter/pkg/types"
)

// TweetParserService implements the TweetReader interface
type TweetParserService struct{}

// NewTweetParser returns a new implementation instance for the TweetReader interface
func NewTweetParser() TweetParser {
	return &TweetParserService{}
}

// FromCSV returns an array of Tweets read from a CSV file
func (p TweetParserService) FromCSV(csv []byte) ([]types.Tweet, error) {
	var tweets []types.Tweet
	if err := csvutil.Unmarshal(csv, &tweets); err != nil {
		fmt.Println("error:", err)
		return tweets, err
	}

	for _, tweet := range tweets {
		fmt.Printf("%+v\n", tweet)
	}
	return tweets, nil
}

// ParseSnippet implements ...
func (p TweetParserService) ParseSnippet(workerID int, tweets <-chan []types.Tweet, decoded chan<- types.Tweet, done chan<- int) {
	done <- workerID
}
