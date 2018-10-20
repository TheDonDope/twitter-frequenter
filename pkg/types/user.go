package types

// User is the struct for the information on the Twitter user
type User struct {
	UserID                 string   `csv:"userid"`
	UserDisplayName        string   `csv:"user_display_name"`
	UserScreenName         string   `csv:"user_screen_name"`
	UserReportedLocation   string   `csv:"user_reported_location"`
	UserProfileDescription string   `csv:"user_profile_description"`
	UserProfileURL         string   `csv:"user_profile_url"`
	FollowerCount          uint64   `csv:"user_follower_count"`
	FollowingCount         uint64   `csv:"user_following_count"`
	AccountCreationDate    UserTime `csv:"account_creation_date"`
	AccountLanguage        string   `csv:"account_language"`
}
