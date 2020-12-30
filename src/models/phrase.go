package models

// Phrase represents a coherent spoken language set of words.
type Phrase struct {
	ID               uint   `json:"id" gorm:"INT; primaryKey"`
	LocaleID         uint   `json:"locale_ID" gorm:"TINYINT NOT NULL"`
	PhraseCategoryID uint   `json:"phrase_category_ID" gorm:" TINYINT NOT NULL"`
	Text             string `json:"name" gorm:"TEXT NOT NULL"`
	Context          string `json:"context" gorm:"TEXT"`
	Creation         string `json:"creation" gorm:"TIMESTAMP"`
}

// CreatePhraseInput represents a coherent spoken language set of words.
type CreatePhraseInput struct {
	ID               uint   `json:"id" binding:"required"`
	LocaleID         uint   `json:"locale_ID" binding:"required"`
	PhraseCategoryID uint   `json:"phrase_category_ID" binding:"required"`
	Text             string `json:"name" binding:"required"`
	Context          string `json:"context" binding:"required"`
	Creation         string `json:"creation" binding:"required"`
}
