package users

import (
	"API/models"
	"API/security"
	"time"
)

type UserService struct {
	User models.User
}

func (receiver *UserService) Save() (models.User, error) {
	pwd, err := security.HashPassword(receiver.User.Password)
	if err != nil {
		return receiver.User, err
	}
	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	user := models.User{
		FirstName: security.SecureString(receiver.User.FirstName),
		LastName:  security.SecureString(receiver.User.LastName),
		UserName:  security.SecureString(security.TrimToLength(receiver.User.UserName, 30)),
		Mail:      security.SecureString(security.TrimToLength(receiver.User.Mail, 252)),
		Password:  pwd,
		Biography: security.SecureString(security.TrimToLength(receiver.User.Biography, 140)),
		Mailing:   security.SecureString(security.TrimToLength(receiver.User.Mailing, 3)),
		Privilege: security.SecureString(security.TrimToLength(receiver.User.Privilege, 3)),
		Points:    uint(receiver.User.Points),
		//TODO: TRANSACTION LOCK FOR CREDIT CHANGE
		Credits:  uint(receiver.User.Credits),
		LocaleID: uint(receiver.User.LocaleID),
		Modified: t,
	}
	if dbError := models.DB.Create(&user); dbError.Error != nil {
		return receiver.User, dbError.Error
	}
	user.Biography = security.RemoveBackticks(user.Biography)
	user.FirstName = security.RemoveBackticks(user.FirstName)
	user.LastName = security.RemoveBackticks(user.LastName)
	user.UserName = security.RemoveBackticks(user.UserName)
	user.Mail = security.RemoveBackticks(user.Mail)
	user.Mailing = security.RemoveBackticks(user.Mailing)
	user.Password = "" // NEVER SEND PWD DATA

	return receiver.User, nil
}
