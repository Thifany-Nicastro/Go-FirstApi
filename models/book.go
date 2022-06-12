package models

type Book struct {
	ID     string `gorm:"primaryKey" binding:"required" json:"id"`
	Title  string `binding:"required" json:"title"`
	Author string `binding:"required" json:"author"`
}
