package models

// WordCategory is a category in which a set of words fall into.
type WordCategory struct {
	ID       uint   `json:"id" gorm:"INT; primaryKey"`
	Name     string `json:"name" gorm:"TEXT NOT NULL"`
	Modified string `json:"modified" gorm:"TIMESTAMP"`
}

// CreateWordCategoryInput is a category in which a set of words fall into.
type CreateWordCategoryInput struct {
	Name string `json:"name" binding:"required"`
}
