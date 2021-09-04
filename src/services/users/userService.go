package services

import (
	"API/models"
	"API/security"
	"errors"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type UserService struct{}

var (
	justPassword = []string{"apple_sub", "facebook_sub", "google_sub"}
	justGoogle   = []string{"password", "apple_sub", "facebook_sub"}
	justFacebook = []string{"password", "apple_sub", "google_sub"}
)

// SignWithHandsApp Just create an user with a jwt token
func (usrService UserService) SignWithHandsApp(receiver models.User) (models.HandsAppJWT, error) {
	return usrService.save(receiver, justPassword...)
}

func (usrService UserService) Login(form models.LoginForm) (models.HandsAppJWT, error) {
	var user models.User
	loginError := errors.New("user or password are incorrect")
	example := models.User{Mail: security.SecureString(form.Credential)}
	err := usrService.findByExample(&user, &example)
	if err != nil {
		return models.HandsAppJWT{}, loginError
	}
	if !security.PasswordMatches(user.Password, form.Password) {
		return models.HandsAppJWT{}, loginError
	}
	return security.CreateJWT(user)
}

// SignWithGoogle Log in a user if exist return his jwt token, otherwise a new user is created with a jwt token
func (usrService UserService) SignWithGoogle(receiver models.User) (models.HandsAppJWT, error) {
	example := models.User{GoogleSub: security.SecureString(receiver.GoogleSub)}
	return usrService.signWith(receiver, example, justGoogle...)
}

// SignWithFacebook Log in a user if exist return his jwt token, otherwise a new user is created with a jwt token
func (usrService UserService) SignWithFacebook(receiver models.User) (models.HandsAppJWT, error) {
	example := models.User{FacebookSub: security.SecureString(receiver.FacebookSub)}
	return usrService.signWith(receiver, example, justFacebook...)
}

// signWith Log in a user if exist return his jwt token, otherwise a new user is created with a jwt token
func (usrService UserService) signWith(receiver, example models.User, omitColumns ...string) (models.HandsAppJWT, error) {
	var user models.User
	err := usrService.findByExample(&user, &example)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return usrService.save(receiver, omitColumns...)
	}
	if err != nil {
		return models.HandsAppJWT{}, err
	}
	return security.CreateJWT(user)
}

func (usrService UserService) findByExample(result, example *models.User) error {
	return models.DB.First(result, example).Error
}

func (usrService UserService) save(receiver models.User, omitColumns ...string) (models.HandsAppJWT, error) {
	var err error
	if receiver.Password, err = security.HashPassword(receiver.Password); err != nil {
		return models.HandsAppJWT{}, err
	}
	user := Safe(receiver)
	user.Modified = time.Now().UTC().Format("2006-01-02 15:04:05")
	user.ID = 0
	user.SubscriberType = ""
	user.Points = 0
	user.Credits = 0
	user.Privilege = ""
	if dbError := models.DB.Omit(omitColumns...).Create(&user); dbError.Error != nil {
		return models.HandsAppJWT{}, dbError.Error
	}

	return security.CreateJWT(user)
}

// Update a user by his id using input.
// If input has a nil filed, that field is not considered updating.
// Returns a new user with backticks.
//
// TODO: move the example to a single file.
// Example
//    // suppose a user with ID = 3, Biography = "boring", FirstName = "Bob"
//
//    newBio := "hello world"
//    usr, _ := Update(4, UpdateUserInput{Mail: &newBio})
//    usr.Bio // "hello world"
//    usr.FirstName // "Bob", since it's nil in the input
func (usrService UserService) Update(id uint64, input models.UpdateUserInput) (models.User, error) {
	user := models.User{ID: uint(id)}

	if err := models.DB.First(&user).Error; err != nil {
		return models.User{}, err
	}

	user = Unsafe(user)

	if input.Password != nil {
		p, err := security.HashPassword(*input.Password)
		if err != nil {
			return models.User{}, err
		}
		input.Password = &p
	}
	newUser := models.User{
		FirstName:   getOrDefault(input.FirstName, user.FirstName).(string),
		LocaleID:    getOrDefault(input.LocaleID, user.LocaleID).(uint),
		LastName:    getOrDefault(input.LastName, user.LastName).(string),
		UserName:    getOrDefault(input.UserName, user.UserName).(string),
		Mail:        getOrDefault(input.Mail, user.Mail).(string),
		Password:    getOrDefault(input.Password, user.Password).(string),
		Biography:   getOrDefault(input.Biography, user.Biography).(string),
		Mailing:     getOrDefault(input.Mailing, user.Mailing).(string),
		GoogleSub:   getOrDefault(input.GoogleSub, user.GoogleSub).(string),
		AppleSub:    getOrDefault(input.AppleSub, user.AppleSub).(string),
		FacebookSub: getOrDefault(input.FacebookSub, user.FacebookSub).(string),
		Picture:     getOrDefault(input.Picture, user.Picture).(string),
		Modified:    time.Now().UTC().Format("2006-01-02 15:04:05"),
	}

	tx := models.DB.
		Model(&user).
		Select("Modified", models.UserInputFields...).
		Updates(Safe(newUser))

	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return user, nil
}

// getOrDefault returns v and defaultValue and v are the same type or if v is nil,
// in other case returns defaultValue.
//
// TODO: move the example to a single file.
// Examples
//     getOrDefault(nil, 23) -> 23
//     getOrDefault(12, 23) -> 12
//     getOrDefault(nil, nil) -> nil
//     getOrDefault('dfs', nil) -> 'dfs'
func getOrDefault(v, defaultValue interface{}) interface{} {
	newValue := reflect.ValueOf(v)
	defaultType := reflect.TypeOf(defaultValue)

	if newValue.Kind() == reflect.Ptr {
		newValue = newValue.Elem()
	}

	// if new value is not nil
	// is better than == because the pointer effects
	if newValue.IsValid() && defaultType == newValue.Type() {
		return newValue.Interface()
	}

	return defaultValue
}

// Unsafe returns a copy for u with all his strings fields without backticks.
func Unsafe(u models.User) models.User {
	return models.User{
		ID:             u.ID,
		FirstName:      security.RemoveBackticks(u.FirstName),
		LastName:       security.RemoveBackticks(u.LastName),
		UserName:       security.RemoveBackticks(u.UserName),
		Mail:           security.RemoveBackticks(u.Mail),
		Password:       u.Password,
		Biography:      security.RemoveBackticks(u.Biography),
		Mailing:        u.Mailing,
		Privilege:      u.Privilege,
		Points:         u.Points,
		Credits:        u.Credits,
		LocaleID:       u.LocaleID,
		Modified:       u.Modified,
		GoogleSub:      security.RemoveBackticks(u.GoogleSub),
		FacebookSub:    security.RemoveBackticks(u.FacebookSub),
		AppleSub:       security.RemoveBackticks(u.AppleSub),
		Picture:        security.RemoveBackticks(u.Picture),
		SubscriberType: u.SubscriberType,
	}
}

// Safe returns a copy for u with all his strings fields secure for sql injection but
// Password, Mailing, Privilege, Modified and SubscriberType.
func Safe(u models.User) models.User {
	return models.User{
		ID:             u.ID,
		FirstName:      security.SecureString(u.FirstName),
		LastName:       security.SecureString(u.LastName),
		UserName:       security.SecureString(security.TrimToLength(u.UserName, 30)),
		Mail:           security.SecureString(security.TrimToLength(u.Mail, 252)),
		Password:       u.Password,
		Biography:      security.SecureString(security.TrimToLength(u.Biography, 140)),
		Mailing:        u.Mailing,
		Privilege:      u.Privilege,
		Points:         u.Points,
		Credits:        u.Credits,
		LocaleID:       u.LocaleID,
		Modified:       u.Modified,
		GoogleSub:      security.SecureString(security.TrimToLength(u.GoogleSub, 68)),
		FacebookSub:    security.SecureString(security.TrimToLength(u.FacebookSub, 68)),
		AppleSub:       security.SecureString(security.TrimToLength(u.AppleSub, 68)),
		Picture:        security.SecureString(security.TrimToLength(u.Picture, 128)),
		SubscriberType: u.SubscriberType,
	}
}
