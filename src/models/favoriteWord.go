package models

// FavoriteWord is a classification of a word by a user.
type FavoriteWord struct {
	WordID   uint   `json:"word_ID" gorm:"INT NOT NULL"`
	UserID   string `json:"user_ID" gorm:"INT NOT NULL"`
	Creation string `json:"creation" gorm:"TIMESTAMP"`
}
