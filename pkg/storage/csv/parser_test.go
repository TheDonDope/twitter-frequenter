package csv

import (
	"strconv"
	"testing"

	"gitlab.com/TheDonDope/gocha/v3/pkg/errors"
)

const exampleTweetCSV = `"tweetid","userid","user_display_name","user_screen_name","user_reported_location","user_profile_description","user_profile_url","follower_count","following_count","account_creation_date","account_language","tweet_language","tweet_text","tweet_time","tweet_client_name","in_reply_to_tweetid","in_reply_to_userid","quoted_tweet_tweetid","is_retweet","retweet_userid","retweet_tweetid","latitude","longitude","quote_count","reply_count","like_count","retweet_count","hashtags","urls","user_mentions","poll_choices"
"877919995476496385","249064136b1c5cb00a705316ab73dd9b53785748ab757f02df7e7a9876906139","249064136b1c5cb00a705316ab73dd9b53785748ab757f02df7e7a9876906139","249064136b1c5cb00a705316ab73dd9b53785748ab757f02df7e7a9876906139","Москва, Россия","Я примерный семьянин!","","132","120","2013-12-07","ru","ru","RT @ruopentwit: ⚡️У НАС НОВОЕ ВИДЕО! Американец: ""Если бы не 27 миллионов русских, я бы сейчас говорил по-немецки""https://t.co/mAcCirn4o1…","2017-06-22 16:03","TweetDeck","","","","true","2572896396","877917212119416832","","","0","0","0","0","[]","[http://ru-open.livejournal.com/374284.html]","[2572896396]",""
"492388766930444288","0974d5dbee4ca9bd6c3b46d62a5cbdbd5c0d86e196b624dbfc7d18cf17b3eab5","0974d5dbee4ca9bd6c3b46d62a5cbdbd5c0d86e196b624dbfc7d18cf17b3eab5","0974d5dbee4ca9bd6c3b46d62a5cbdbd5c0d86e196b624dbfc7d18cf17b3eab5","Россия","Телефонист .Изучение истории   Игра в любительском театре  -  Воздушные змеи ,","","74","8","2014-03-15","en","ru","Серебром отколоколило http://t.co/Jaa4v4IFpM","2014-07-24 19:20","generationπ","","","","false","","","","","0","0","0","0","","[http://pyypilg33.livejournal.com/11069.html]","",""`

const exampleUserCSV = `"userid","user_display_name","user_screen_name","user_reported_location","user_profile_description","user_profile_url","follower_count","following_count","account_creation_date","account_language"
"94c383089f0dd9993020276bd01113ecb5935ad860bfa61e6079e7d548577f76","94c383089f0dd9993020276bd01113ecb5935ad860bfa61e6079e7d548577f76","94c383089f0dd9993020276bd01113ecb5935ad860bfa61e6079e7d548577f76","USA","Conservative. Politics. Pro-Life. Writer.  #GodBlessAmerica #WakeUpAmerica #InGodWeTrust","","1541","1676","2014-06-13","en"
"799bb58d3c064d3884fdc4604c74068ff059b60c0b34642621c6f638fb436f6d","799bb58d3c064d3884fdc4604c74068ff059b60c0b34642621c6f638fb436f6d","799bb58d3c064d3884fdc4604c74068ff059b60c0b34642621c6f638fb436f6d","Раша","На воре и шапка горит","","299","324","2013-12-22","en"`

func TestTweetsFromCSV(t *testing.T) {
	// given
	underTest := NewTweetParser()

	// when
	actualTweets, err := underTest.FromCSV([]byte(exampleTweetCSV))

	// then
	if err != nil {
		t.Errorf(errors.Fail("File", "<nil>", err.Error()))
	}

	expectedSize := 2
	actualSize := len(actualTweets)
	if expectedSize != actualSize {
		t.Errorf(errors.Fail("Size of result", strconv.Itoa(expectedSize), strconv.Itoa(actualSize)))
	}
}

func TestUsersFromCSV(t *testing.T) {
	// given
	underTest := NewUserParser()

	// when
	actualUsers, err := underTest.FromCSV([]byte(exampleUserCSV))

	// then
	if err != nil {
		t.Errorf(errors.Fail("File", "<nil>", err.Error()))
	}

	expectedSize := 2
	actualSize := len(actualUsers)
	if expectedSize != actualSize {
		t.Errorf(errors.Fail("Size of result", strconv.Itoa(expectedSize), strconv.Itoa(actualSize)))
	}
}
