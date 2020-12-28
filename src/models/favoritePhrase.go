package models

// FavoritePhrase is a classification of a phrase by a user.
type FavoritePhrase struct {
	PhraseID uint   `json:"phrase_ID"`
	UserID   string `json:"user_ID" gorm:"not null"`
	Creation string `json:"creation"`
}
