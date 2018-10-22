package csv

// Tweet is the struct for the information on the tweet
type Tweet struct {
	TweetID                string    `csv:"tweetid"`
	UserID                 string    `csv:"userid"`
	UserDisplayName        string    `csv:"user_display_name"`
	UserScreenName         string    `csv:"user_screen_name"`
	UserReportedLocation   string    `csv:"user_reported_location"`
	UserProfileDescription string    `csv:"user_profile_description"`
	UserProfileURL         string    `csv:"user_profile_url"`
	FollowerCount          uint64    `csv:"user_follower_count"`
	FollowingCount         uint64    `csv:"user_following_count"`
	AccountCreationDate    UserTime  `csv:"account_creation_date"`
	AccountLanguage        string    `csv:"account_language"`
	TweetLanguage          string    `csv:"tweet_language"`
	TweetText              string    `csv:"tweet_text"`
	TweetTime              TweetTime `csv:"tweet_time"`
	TweetClientName        string    `csv:"tweet_client_name"`
	InReplyToTweetID       string    `csv:"in_reply_to_tweetid"`
	InReplyToUserID        string    `csv:"in_reply_to_userid"`
	QuotedTweetTweetID     string    `csv:"quoted_tweet_tweetid"`
	IsRetweet              string    `csv:"is_retweet"`
	RetweetUserID          string    `csv:"retweet_user_id"`
	RetweetTweetID         string    `csv:"retweet_tweet_id"`
	Latitude               string    `csv:"latitude"`
	Longitude              string    `csv:"longitude"`
	QuoteCount             string    `csv:"quote_count"`
	ReplyCount             string    `csv:"reply_count"`
	LikeCount              string    `csv:"like_count"`
	RetweetCount           string    `csv:"retweet_count"`
	Hashtags               string    `csv:"hashtags"`
	URLs                   string    `csv:"urls"`
	UserMentions           string    `csv:"user_mentions"`
	PollChoices            string    `csv:"poll_choices"`
}
