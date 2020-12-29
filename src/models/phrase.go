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
