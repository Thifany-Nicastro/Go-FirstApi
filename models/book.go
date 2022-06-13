package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID     string `gorm:"primaryKey;unique;not null" json:"id"`
	Title  string `gorm:"unique;not null" json:"title"`
	Author string `gorm:"not null" json:"author"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewString()

	return
}
