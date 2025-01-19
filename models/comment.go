package models

// Comment struct untuk representasi komentar
type Comment struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	UserID  uint   `json:"userId"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}
