package services

import (
	. "API/models"
	"fmt"
	"testing"
	"time"
)

var (
	userService         = UserService{}
	columnsWithPassword = []string{"google_sub", "apple_sub", "facebook_sub"}
)

func init() {
	ConnectDatabase()
}

// Test the normal use case whe using password and mail
func TestUserService_save(t *testing.T) {
	uniqueSalt := time.Now().UTC().Unix()
	token, err := userService.save(User{
		LocaleID:  1,
		Mail:      fmt.Sprintf("example%v@example.com", uniqueSalt),
		UserName:  fmt.Sprintf("exampleUser%v", uniqueSalt),
		FirstName: fmt.Sprintf("exampleFstName%v", uniqueSalt),
		LastName:  fmt.Sprintf("exampleLstName%v", uniqueSalt),
		Password:  "super_secret_pass",
	}, columnsWithPassword...)

	if err != nil {
		t.Fatalf(`expect no error but got: %v`, err.Error())
	}

	if token.Token == "" {
		t.Fail()
	}
}

func TestUserService_saveWithWrongParams(t *testing.T) {
	wrongUsers := []User{
		{},                // without fields
		{LocaleID: 32132}, // unknown locale
		{LocaleID: 1, Mail: "example@example.com"},                                                             // without username
		{LocaleID: 1, Mail: "example@example.com", UserName: "dummyUser"},                                      // without password
		{LocaleID: 1, GoogleSub: "", UserName: "dummyUser"},                                                    // empty google id
		{LocaleID: 1, GoogleSub: "asasasa", AppleSub: "dsadw21", FacebookSub: "fdsfds", UserName: "dummyUser"}, // multiple sign up methods in a row
	}
	for _, wrongUser := range wrongUsers {
		if _, err := userService.save(wrongUser); err == nil {
			t.Fail()
		}
	}
}
