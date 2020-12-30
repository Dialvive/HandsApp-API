package models

// FavoritePhrase is a classification of a phrase by a user.
type FavoritePhrase struct {
	PhraseID uint   `json:"phrase_ID" gorm:"INT NOT NULL"`
	UserID   string `json:"user_ID" gorm:"INT NOT NULL"`
	Modified string `json:"modified" gorm:"TIMESTAMP"`
}

// CreateFavoritePhraseInput is a classification of a phrase by a user.
type CreateFavoritePhraseInput struct {
	PhraseID uint   `json:"phrase_ID" binding:"required"`
	UserID   string `json:"user_ID" binding:"required"`
}
