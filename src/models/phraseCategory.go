package models

// PhraseCategory is a category in which a set of phrases fall into.
type PhraseCategory struct {
	ID       uint   `json:"id" gorm:"INT; primaryKey"`
	Name     string `json:"name" gorm:"TEXT NOT NULL"`
	Modified string `json:"modified" gorm:"TIMESTAMP"`
}

// CreatePhraseCategoryInput is a category in which a set of phrases fall into.
type CreatePhraseCategoryInput struct {
	Name string `json:"name" binding:"required"`
}
