package services

import (
	"API/models"
	"API/security"
	"errors"
	"gorm.io/gorm"
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
	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	hashPassword, err := security.HashPassword(receiver.Password)
	if err != nil {
		return models.HandsAppJWT{}, err
	}
	user := models.User{
		ID:             0,
		FirstName:      security.SecureString(receiver.FirstName),
		LastName:       security.SecureString(receiver.LastName),
		UserName:       security.SecureString(security.TrimToLength(receiver.UserName, 30)),
		Mail:           security.SecureString(security.TrimToLength(receiver.Mail, 252)),
		Password:       hashPassword,
		Biography:      security.SecureString(security.TrimToLength(receiver.Biography, 140)),
		Mailing:        receiver.Mailing,
		Privilege:      receiver.Privilege,
		Points:         receiver.Points,
		Credits:        receiver.Credits,
		LocaleID:       receiver.LocaleID,
		Modified:       t,
		GoogleSub:      security.SecureString(security.TrimToLength(receiver.GoogleSub, 68)),
		FacebookSub:    security.SecureString(security.TrimToLength(receiver.FacebookSub, 68)),
		AppleSub:       security.SecureString(security.TrimToLength(receiver.AppleSub, 68)),
		Picture:        security.SecureString(security.TrimToLength(receiver.Picture, 128)),
		SubscriberType: receiver.SubscriberType,
	}

	if dbError := models.DB.Omit(omitColumns...).Create(&user); dbError.Error != nil {
		return models.HandsAppJWT{}, dbError.Error
	}

	return security.CreateJWT(user)
}
