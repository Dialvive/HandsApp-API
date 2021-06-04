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
	user := models.User{
		FirstName: security.SecureString(receiver.FirstName),
		LastName:  security.SecureString(receiver.LastName),
		UserName:  security.SecureString(security.TrimToLength(receiver.UserName, 30)),
		Mail:      security.SecureString(security.TrimToLength(receiver.Mail, 252)),
		Biography: security.SecureString(security.TrimToLength(receiver.Biography, 140)),
		Mailing:   security.SecureString(security.TrimToLength(receiver.Mailing, 12)),
		Privilege: security.SecureString(security.TrimToLength(receiver.Privilege, 10)),
		Points:    receiver.Points,
		Credits:   receiver.Credits,
		LocaleID:  receiver.LocaleID,
		GoogleSub: security.SecureString(security.TrimToLength(receiver.GoogleSub, 68)),
		Modified:  t,
		Picture:   security.SecureString(security.TrimToLength(receiver.Picture, 128)),
	}

	pwd, err := security.HashPassword(receiver.Password)
	if err != nil {
		return "", err
	}
	user.Password = pwd

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
	form.Identifier = security.SecureString(form.Identifier)
	if res := models.DB.Where(&models.User{Mail: form.Identifier}).Or(&models.User{UserName: form.Identifier}).First(&user); res.Error != nil {
		return "", loginError
	}
	if !security.PasswordMatches(user.Password, form.Password) {
		return "", loginError
	}
	return security.CreateJWT(user), nil
}
