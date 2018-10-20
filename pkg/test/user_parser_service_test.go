package test

import (
	"strconv"
	"testing"

	"gitlab.com/TheDonDope/twitter-frequenter/pkg/api"
	ourErrors "gitlab.com/TheDonDope/twitter-frequenter/pkg/util/errors"
)

const exampleUserCSV = `"userid","user_display_name","user_screen_name","user_reported_location","user_profile_description","user_profile_url","follower_count","following_count","account_creation_date","account_language"
"94c383089f0dd9993020276bd01113ecb5935ad860bfa61e6079e7d548577f76","94c383089f0dd9993020276bd01113ecb5935ad860bfa61e6079e7d548577f76","94c383089f0dd9993020276bd01113ecb5935ad860bfa61e6079e7d548577f76","USA","Conservative. Politics. Pro-Life. Writer.  #GodBlessAmerica #WakeUpAmerica #InGodWeTrust","","1541","1676","2014-06-13","en"
"799bb58d3c064d3884fdc4604c74068ff059b60c0b34642621c6f638fb436f6d","799bb58d3c064d3884fdc4604c74068ff059b60c0b34642621c6f638fb436f6d","799bb58d3c064d3884fdc4604c74068ff059b60c0b34642621c6f638fb436f6d","Раша","На воре и шапка горит","","299","324","2013-12-22","en"`

func TestUsersFromCSV(t *testing.T) {
	// given
	underTest := api.NewUserParser()

	// when
	actualUsers, err := underTest.FromCSV([]byte(exampleUserCSV))

	// then
	if err != nil {
		t.Errorf(ourErrors.GetFormattedFailMessage("File", "<nil>", err.Error()))
	}

	expectedSize := 2
	actualSize := len(actualUsers)
	if expectedSize != actualSize {
		t.Errorf(ourErrors.GetFormattedFailMessage("Size of result", strconv.Itoa(expectedSize), strconv.Itoa(actualSize)))
	}
}
