package csv

import "time"

// UserTime is a custom Time struct to be able to override the UnmarshalCSV method
// Format: "2017-06-22"
type UserTime struct {
	time.Time
}

// UnmarshalCSV reads the given data bytes and formats it as a UserTime object
func (t *UserTime) UnmarshalCSV(data []byte) error {
	tt, err := time.Parse("2006-01-02", string(data))
	if err != nil {
		return err
	}
	*t = UserTime{Time: tt}
	return nil
}

// TweetTime is a custom Time struct to be able to override the UnmarshalCSV method
// Format: "2017-06-22 16:03"
type TweetTime struct {
	time.Time
}

// UnmarshalCSV reads the given data bytes and formats it as a TweetTime object
func (t *TweetTime) UnmarshalCSV(data []byte) error {
	tt, err := time.Parse("2006-01-02 15:04", string(data))
	if err != nil {
		return err
	}
	*t = TweetTime{Time: tt}
	return nil
}
