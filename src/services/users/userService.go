package services

import (
	"API/models"
	"API/security"
	"time"
)

type UserService struct {
	User models.User
}

func (usrService *UserService) Save(receiver models.User, omitColumns ...string) (models.User, error) {
	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	user := models.User{
		FirstName: security.SecureString(receiver.FirstName),
		LastName:  security.SecureString(receiver.LastName),
		UserName:  security.SecureString(security.TrimToLength(receiver.UserName, 30)),
		Mail:      security.SecureString(security.TrimToLength(receiver.Mail, 252)),
		Biography: security.SecureString(security.TrimToLength(receiver.Biography, 140)),
		Mailing:   security.SecureString(security.TrimToLength(receiver.Mailing, 3)),
		Privilege: security.SecureString(security.TrimToLength(receiver.Privilege, 3)),
		Points:    receiver.Points,
		Credits:   receiver.Credits,
		LocaleID:  receiver.LocaleID,
		GoogleSub: security.SecureString(receiver.GoogleSub),
		Modified:  t,
	}

	pwd, err := security.HashPassword(receiver.Password)
	if err != nil {
		return receiver, err
	}
	user.Password = pwd

	if dbError := models.DB.Omit(omitColumns...).Create(&user); dbError.Error != nil {
		return receiver, dbError.Error
	}

	user.Biography = security.RemoveBackticks(user.Biography)
	user.FirstName = security.RemoveBackticks(user.FirstName)
	user.LastName = security.RemoveBackticks(user.LastName)
	user.UserName = security.RemoveBackticks(user.UserName)
	user.Mail = security.RemoveBackticks(user.Mail)
	user.Mailing = security.RemoveBackticks(user.Mailing)
	user.Password = "" // NEVER SEND PWD DATA

	return user, nil
}
