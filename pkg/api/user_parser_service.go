package api

import (
	"fmt"

	"github.com/jszwec/csvutil"
	"gitlab.com/TheDonDope/twitter-frequenter/pkg/types"
)

// UserParserService implements the UserReader interface
type UserParserService struct{}

// NewUserParser returns a new implementation instance for the UserReader interface
func NewUserParser() UserParser {
	return &UserParserService{}
}

// FromCSV returns an array of Users read from a CSV file
func (p UserParserService) FromCSV(csv []byte) ([]types.User, error) {
	var users []types.User
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
func (p UserParserService) ParseSnippet(workerID int, users <-chan []types.User, decoded chan<- types.User, done chan<- int) {
	proc := 0 // How many users did we process?
	for userArray := range users {
		for _, user := range userArray {
			proc++
			decoded <- user
		}
	}
	done <- workerID
}
