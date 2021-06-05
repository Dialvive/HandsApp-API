package services

import (
	"API/models"
	"API/security"
	"errors"
	"time"
)

type UserService struct{}

func (usrService *UserService) Save(receiver models.User, omitColumns ...string) (string, error) {
	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	hashPassword, err := security.HashPassword(receiver.Password)
	if err != nil {
		return "", err
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
		return "", dbError.Error
	}

	user.Biography = security.RemoveBackticks(user.Biography)
	user.FirstName = security.RemoveBackticks(user.FirstName)
	user.LastName = security.RemoveBackticks(user.LastName)
	user.UserName = security.RemoveBackticks(user.UserName)
	user.Mail = security.RemoveBackticks(user.Mail)
	user.Mailing = security.RemoveBackticks(user.Mailing)
	user.Privilege = security.RemoveBackticks(user.Privilege)
	user.Password = "" // NEVER SEND PWD DATA

	signedString := security.CreateJWT(user)
	return signedString, nil
}

func (usrService UserService) Login(form models.LoginForm) (string, error) {
	var user models.User
	loginError := errors.New("user or password are incorrect")
	form.Credential = security.SecureString(form.Credential)
	if res := models.DB.Where(&models.User{Mail: form.Credential}).Or(&models.User{UserName: form.Credential}).First(&user); res.Error != nil {
		return "", loginError
	}
	if !security.PasswordMatches(user.Password, form.Password) {
		return "", loginError
	}
	return security.CreateJWT(user), nil
}

func (usrService UserService) LoginWithGoogle(form models.LoginForm) (string, error) {
	form.Credential = security.SecureString(form.Credential)
	var user models.User
	if res := models.DB.Where(&models.User{GoogleSub: form.Credential}).First(&user); res.Error != nil {
		return "", res.Error
	}
	return security.CreateJWT(user), nil
}
