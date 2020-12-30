package models

// FavoriteWord is a classification of a word by a user.
type FavoriteWord struct {
	WordID   uint   `json:"word_ID" gorm:"INT NOT NULL"`
	UserID   string `json:"user_ID" gorm:"INT NOT NULL"`
	Modified string `json:"modified" gorm:"TIMESTAMP"`
}

// CreateFavoriteWordInput is a classification of a word by a user.
type CreateFavoriteWordInput struct {
	WordID uint   `json:"word_ID" binding:"required"`
	UserID string `json:"user_ID" binding:"required"`
}
