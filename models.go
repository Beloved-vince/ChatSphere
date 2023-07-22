package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID        uint      `gorm:"primary_key;auto_increment"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Username  string    `gorm:"not null;unique"`
	Email     *string   `gorm:"not null;unique"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Room struct {
	ID        uint      `gorm:"primary_key;auto_increment"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Message struct {
	ID        uint      `gorm:"primary_key;auto_increment"`
	UserID    uint      `gorm:"not null"`
	RoomID    uint      `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
