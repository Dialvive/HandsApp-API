package services

import (
	. "API/models"
	"API/security"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"math/rand"
	"reflect"
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

func TestUserService_Update(t *testing.T) {
	randomData := dummyUserWithPassword()

	testCases := []struct {
		name    string
		changes UpdateUserInput
	}{
		{"update Mail", UpdateUserInput{Mail: &randomData.UserName}},
		{"update FirstName", UpdateUserInput{FirstName: &randomData.FirstName}},
		{"update LastName", UpdateUserInput{LastName: &randomData.LastName}},
		{"update UserName", UpdateUserInput{UserName: &randomData.UserName}},
		{"update Password", UpdateUserInput{Password: &randomData.Password}},
		{"update Biography", UpdateUserInput{Biography: &randomData.Biography}},
		{"update Mailing", UpdateUserInput{Mailing: &randomData.Mailing}},
		{"update LocaleID", UpdateUserInput{LocaleID: &randomData.LocaleID}},
		{"update GoogleSub", UpdateUserInput{GoogleSub: &randomData.GoogleSub}},
		{"update FacebookSub", UpdateUserInput{FacebookSub: &randomData.FacebookSub}},
		{"update AppleSub", UpdateUserInput{AppleSub: &randomData.AppleSub}},
		{"update Picture", UpdateUserInput{Picture: &randomData.Picture}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			oldUser := dummyUserWithPassword()
			var newUser User

			_, err := userService.SignWithHandsApp(oldUser)
			if err != nil {
				t.Fatal("create user got:", err.Error())
			}

			tx := DB.Take(&newUser, "user_name = ?", security.SecureString(oldUser.UserName))
			if tx.Error != nil {
				t.Fatal("create user got:", err.Error())
			}

			newUser, err = userService.Update(uint64(newUser.ID), tc.changes)
			if err != nil {
				t.Fatal("update user got: ", err.Error())
			}

			DB.Take(&newUser, "user_name = ?", security.SecureString(oldUser.UserName))

			newUser = Unsafe(newUser)

			baseFieldsToHide := []string{
				"ID",
				"Mailing",
				"Privilege",
				"Password",
				"SubscriberType",
				"Modified",
				"GoogleSub",
				"FacebookSub",
				"AppleSub",
			}
			ignoreFields := cmpopts.IgnoreFields(User{}, append(baseFieldsToHide, getNonNilFields(tc.changes)...)...)

			equal := cmp.Equal(newUser, oldUser, ignoreFields)
			if !equal {
				t.Fatalf(
					"old user expect to not be equal to new user, diff:\n %v",
					cmp.Diff(newUser, oldUser, ignoreFields),
				)
			}
		})
	}
}

func getNonNilFields(u UpdateUserInput) []string {
	uType := reflect.ValueOf(u)
	var fields []string
	for i := 0; i < uType.NumField(); i++ {
		field := uType.Field(i)
		if !field.IsNil() {
			fieldName := uType.Type().Field(i).Name
			fields = append(fields, fieldName)
		}
	}
	return fields
}

func TestUserService_UpdatePassword(t *testing.T) {
	dummyUser := dummyUserWithPassword()
	oldPlainPassword := dummyUser.Password
	newPassword := "hello"
	_, err := userService.SignWithHandsApp(dummyUser)

	if err != nil {
		t.Fatal("create user got:", err.Error())
	}

	tx := DB.Take(&dummyUser, "user_name = ?", security.SecureString(dummyUser.UserName))

	if tx.Error != nil {
		t.Fatal("create user got:", err.Error())
	}

	newUser, err := userService.Update(uint64(dummyUser.ID), UpdateUserInput{Password: &newPassword})

	if err != nil {
		t.Fatal("update user got: ", err.Error())
	}

	if dummyUser.Password == newUser.Password {
		t.Fatalf("%v expected not to be equal to %v", dummyUser.Password, newUser.Password)
	}

	if security.PasswordMatches(newUser.Password, oldPlainPassword) {
		t.Fatalf("shouldn't match with the old plain password %v", oldPlainPassword)
	}

	// retrieve the user again because the new hashed password was empty
	DB.Where("user_name = ?", dummyUser.UserName).First(&newUser)

	if !security.PasswordMatches(newUser.Password, newPassword) {
		t.Fatalf("should match with the new plain password %v", oldPlainPassword)
	}
}

func dummyUserWithPassword() User {
	rand.Seed(time.Now().UnixNano())
	uniqueSalt := rand.Uint32()
	return User{
		FirstName: fmt.Sprint("exampleFstName", uniqueSalt),
		LastName:  fmt.Sprint("exampleLstName", uniqueSalt),
		UserName:  fmt.Sprint("exampleUser", uniqueSalt),
		Mail:      fmt.Sprint("example", uniqueSalt, "@example.com"),
		Password:  "super_secret_pass",
		Biography: fmt.Sprint("bio", uniqueSalt),
		LocaleID:  1,
		Picture:   fmt.Sprint("https://example/", uniqueSalt),
	}
}
