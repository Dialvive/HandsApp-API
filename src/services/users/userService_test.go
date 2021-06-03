package services

import (
	. "API/models"
	"testing"
)

var (
	userService         = UserService{}
	columnsWithPassword = []string{"google_sub", "apple_sub", "facebook_sub"}
)

func init() {
	ConnectDatabase()
}

// Test the normal use case whe using password and mail
func TestUserService_Save(t *testing.T) {
	user, err := userService.Save(User{
		LocaleID: 1,
		Mail:     "example@example.com",
		UserName: "exampleUser",
		Password: "super_secret_pass",
	}, columnsWithPassword...)

	if err != nil {
		t.Fatalf(`expect no error but got: %v`, err.Error())
	}

	if user.ID == 0 {
		t.Fail()
	}
}

func TestUserService_SaveWithWrongParams(t *testing.T) {
	wrongUsers := []User{
		{},                // without fields
		{LocaleID: 32132}, // unknown locale
		{LocaleID: 1, Mail: "example@example.com"},                                                             // without username
		{LocaleID: 1, Mail: "example@example.com", UserName: "dummyUser"},                                      // without password
		{LocaleID: 1, GoogleSub: "", UserName: "dummyUser"},                                                    // empty google id
		{LocaleID: 1, GoogleSub: "asasasa", AppleSub: "dsadw21", FacebookSub: "fdsfds", UserName: "dummyUser"}, // multiple sign up methods in a row
	}
	for _, wrongUser := range wrongUsers {
		if _, err := userService.Save(wrongUser); err == nil {
			t.Fail()
		}
	}
}
