package model

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Name         string    `gorm:"size:100;not null"`
	Email        string    `gorm:"size:100;unique;not null"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
