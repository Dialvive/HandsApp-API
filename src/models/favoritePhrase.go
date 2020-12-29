package models

// FavoritePhrase is a classification of a phrase by a user.
type FavoritePhrase struct {
	PhraseID uint   `json:"phrase_ID" gorm:"INT NOT NULL"`
	UserID   string `json:"user_ID" gorm:"INT NOT NULL"`
	Creation string `json:"creation" gorm:"TIMESTAMP"`
}
