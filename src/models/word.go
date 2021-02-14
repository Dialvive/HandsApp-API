package models

import "github.com/dgrijalva/jwt-go"

// Word represents a coherent spoken language set of words.
type Word struct {
	ID             uint   `json:"ID" gorm:"column:ID;type:INT NOT NULL;primaryKey"`
	LocaleID       uint   `json:"locale_ID" gorm:"column:locale_ID;type:TINYINT NOT NULL"`
	WordCategoryID uint   `json:"word_category_ID" gorm:"column:word_category_ID;type:TINYINT NOT NULL"`
	TextDe         string `json:"text_de" gorm:"column:text_de;type:TEXT"`
	TextEs         string `json:"text_es" gorm:"column:text_es;type:TEXT"`
	TextEn         string `json:"text_en" gorm:"column:text_en;type:TEXT"`
	TextFr         string `json:"text_fr" gorm:"column:text_fr;type:TEXT"`
	TextIt         string `json:"text_it" gorm:"column:text_it;type:TEXT"`
	TextPt         string `json:"text_pt" gorm:"column:text_pt;type:TEXT"`
	ContextDe      string `json:"context_de" gorm:"column:context_de;type:TEXT"`
	ContextEs      string `json:"context_es" gorm:"column:context_es;type:TEXT"`
	ContextEn      string `json:"context_en" gorm:"column:context_en;type:TEXT"`
	ContextFr      string `json:"context_fr" gorm:"column:context_fr;type:TEXT"`
	ContextIt      string `json:"context_it" gorm:"column:context_it;type:TEXT"`
	ContextPt      string `json:"context_pt" gorm:"column:context_pt;type:TEXT"`
	DefinitionDe   string `json:"definition_de" gorm:"column:definition_de;type:TEXT"`
	DefinitionEs   string `json:"definition_es" gorm:"column:definition_es;type:TEXT"`
	DefinitionEn   string `json:"definition_en" gorm:"column:definition_en;type:TEXT"`
	DefinitionFr   string `json:"definition_fr" gorm:"column:definition_fr;type:TEXT"`
	DefinitionIt   string `json:"definition_it" gorm:"column:definition_it;type:TEXT"`
	DefinitionPt   string `json:"definition_pt" gorm:"column:definition_pt;type:TEXT"`
	Modified       string `json:"modified" gorm:"column:modified;type:TIMESTAMP"`
}

// CreateWordInput represents a coherent spoken word.
type CreateWordInput struct {
	LocaleID       uint   `json:"locale_ID" binding:"required"`
	WordCategoryID uint   `json:"word_category_ID" binding:"required"`
	TextDe         string `json:"text_de"`
	TextEs         string `json:"text_es"`
	TextEn         string `json:"text_en"`
	TextFr         string `json:"text_fr"`
	TextIt         string `json:"text_it"`
	TextPt         string `json:"text_pt"`
	ContextDe      string `json:"context_de"`
	ContextEs      string `json:"context_es"`
	ContextEn      string `json:"context_en"`
	ContextFr      string `json:"context_fr"`
	ContextIt      string `json:"context_it"`
	ContextPt      string `json:"context_pt"`
	DefinitionDe   string `json:"definition_de"`
	DefinitionEs   string `json:"definition_es"`
	DefinitionEn   string `json:"definition_en"`
	DefinitionFr   string `json:"definition_fr"`
	DefinitionIt   string `json:"definition_it"`
	DefinitionPt   string `json:"definition_pt"`
}

// UpdateWordInput represents a coherent spoken word.
type UpdateWordInput struct {
	LocaleID       uint   `json:"locale_ID"`
	WordCategoryID uint   `json:"word_category_ID"`
	TextDe         string `json:"text_de"`
	TextEs         string `json:"text_es"`
	TextEn         string `json:"text_en"`
	TextFr         string `json:"text_fr"`
	TextIt         string `json:"text_it"`
	TextPt         string `json:"text_pt"`
	ContextDe      string `json:"context_de"`
	ContextEs      string `json:"context_es"`
	ContextEn      string `json:"context_en"`
	ContextFr      string `json:"context_fr"`
	ContextIt      string `json:"context_it"`
	ContextPt      string `json:"context_pt"`
	DefinitionDe   string `json:"definition_de"`
	DefinitionEs   string `json:"definition_es"`
	DefinitionEn   string `json:"definition_en"`
	DefinitionFr   string `json:"definition_fr"`
	DefinitionIt   string `json:"definition_it"`
	DefinitionPt   string `json:"definition_pt"`
}

// WordClaim is  a claim that cointains Word as Data.
type WordClaim struct {
	Data Word `json:"data" binding:"required"`
	jwt.StandardClaims
}
