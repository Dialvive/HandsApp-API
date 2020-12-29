package models

// WordCategory is a category in which a set of words fall into.
type WordCategory struct {
	ID       uint   `json:"id" gorm:"INT; primaryKey"`
	Name     string `json:"name" gorm:"TEXT NOT NULL"`
	Creation string `json:"creation" gorm:"TIMESTAMP"`
}
