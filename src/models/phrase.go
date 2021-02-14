package models

import "github.com/dgrijalva/jwt-go"

// Phrase represents a coherent spoken language set of words.
type Phrase struct {
	ID               uint   `json:"ID" gorm:"column:ID;type:INT NOT NULL;primaryKey"`
	LocaleID         uint   `json:"locale_ID" gorm:"column:locale_ID;type:TINYINT NOT NULL"`
	PhraseCategoryID uint   `json:"phrase_category_ID" gorm:"column:phrase_category_ID;type:TINYINT NOT NULL"`
	Text             string `json:"text" gorm:"column:text;type:TEXT"`
	TextDe           string `json:"text_de" gorm:"column:text_de;type:TEXT"`
	TextEs           string `json:"text_es" gorm:"column:text_es;type:TEXT"`
	TextEn           string `json:"text_en" gorm:"column:text_en;type:TEXT"`
	TextFr           string `json:"text_fr" gorm:"column:text_fr;type:TEXT"`
	TextIt           string `json:"text_it" gorm:"column:text_it;type:TEXT"`
	TextPt           string `json:"text_pt" gorm:"column:text_pt;type:TEXT"`
	ContextDe        string `json:"context_de" gorm:"column:context_de;type:TEXT"`
	ContextEs        string `json:"context_es" gorm:"column:context_es;type:TEXT"`
	ContextEn        string `json:"context_en" gorm:"column:context_en;type:TEXT"`
	ContextFr        string `json:"context_fr" gorm:"column:context_fr;type:TEXT"`
	ContextIt        string `json:"context_it" gorm:"column:context_it;type:TEXT"`
	ContextPt        string `json:"context_pt" gorm:"column:context_pt;type:TEXT"`
	Modified         string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreatePhraseInput represents a coherent spoken language set of words.
type CreatePhraseInput struct {
	LocaleID         uint   `json:"locale_ID" binding:"required"`
	PhraseCategoryID uint   `json:"phrase_category_ID" binding:"required"`
	Text             string `json:"text" gorm:"column:text;type:TEXT" binding:"required"`
	TextDe           string `json:"text_de"`
	TextEs           string `json:"text_es"`
	TextEn           string `json:"text_en"`
	TextFr           string `json:"text_fr"`
	TextIt           string `json:"text_it"`
	TextPt           string `json:"text_pt"`
	ContextDe        string `json:"context_de"`
	ContextEs        string `json:"context_es"`
	ContextEn        string `json:"context_en"`
	ContextFr        string `json:"context_fr"`
	ContextIt        string `json:"context_it"`
	ContextPt        string `json:"context_pt"`
}

// UpdatePhraseInput represents a coherent spoken language set of words.
type UpdatePhraseInput struct {
	LocaleID         uint   `json:"locale_ID"`
	PhraseCategoryID uint   `json:"phrase_category_ID"`
	Text             string `json:"text" gorm:"column:text;type:TEXT"`
	TextDe           string `json:"text_de"`
	TextEs           string `json:"text_es"`
	TextEn           string `json:"text_en"`
	TextFr           string `json:"text_fr"`
	TextIt           string `json:"text_it"`
	TextPt           string `json:"text_pt"`
	ContextDe        string `json:"context_de"`
	ContextEs        string `json:"context_es"`
	ContextEn        string `json:"context_en"`
	ContextFr        string `json:"context_fr"`
	ContextIt        string `json:"context_it"`
	ContextPt        string `json:"context_pt"`
}

// PhraseClaim is  a claim that cointains Phrase as Data.
type PhraseClaim struct {
	Data Phrase `json:"data" binding:"required"`
	jwt.StandardClaims
}
