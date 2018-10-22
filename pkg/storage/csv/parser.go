package csv

import (
	"fmt"

	"github.com/jszwec/csvutil"
)

// TweetParser defines methods to read tweet data from a source CSV
type TweetParser interface {
	// FromCSV returns an array of Tweets read from a CSV file
	FromCSV([]byte) ([]Tweet, error)

	// ParseSnippet parses ...
	ParseSnippet(int, <-chan []Tweet, chan<- Tweet, chan<- int)
}

// UserParser defines methods to read twitter user data from a source CSV
type UserParser interface {
	// FromCSV returns an array of TwitterUsers read from a CSV file
	FromCSV([]byte) ([]User, error)

	// ParseSnippet parses ...
	ParseSnippet(int, <-chan []User, chan<- User, chan<- int)
}

type tweetParser struct{}

type userParser struct{}

// NewTweetParser returns a new implementation instance for the TweetParser interface
func NewTweetParser() TweetParser {
	return &tweetParser{}
}

// NewUserParser returns a new implementation instance for the UserParser interface
func NewUserParser() UserParser {
	return &userParser{}
}

// FromCSV returns an array of Tweets read from a CSV file
func (t *tweetParser) FromCSV(csv []byte) ([]Tweet, error) {
	var tweets []Tweet
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
func (t *tweetParser) ParseSnippet(workerID int, tweets <-chan []Tweet, decoded chan<- Tweet, done chan<- int) {
	done <- workerID
}

// FromCSV returns an array of Users read from a CSV file
func (u *userParser) FromCSV(csv []byte) ([]User, error) {
	var users []User
	if err := csvutil.Unmarshal(csv, &users); err != nil {
		fmt.Println("error:", err)
		return users, err
	}

	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
	return users, nil
}

// ParseSnippet implements ...
func (u *userParser) ParseSnippet(workerID int, users <-chan []User, decoded chan<- User, done chan<- int) {
	proc := 0 // How many users did we process?
	for userArray := range users {
		for _, user := range userArray {
			proc++
			decoded <- user
		}
	}
	done <- workerID
}
