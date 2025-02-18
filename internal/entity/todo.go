package entity

import "time"

type Todo struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null"`
	Title     string    `gorm:"size:255;not null"`
	Completed bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
