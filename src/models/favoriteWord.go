package models

// FavoriteWord is a classification of a word by a user.
type FavoriteWord struct {
	WordID   uint   `json:"word_ID"`
	UserID   string `json:"user_ID" gorm:"not null"`
	Creation string `json:"creation"`
}
