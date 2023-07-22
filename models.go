package main

import "github.com/jinzhu/gorm"
import "time"

type User struct {
	gorm.Model
	ID        uint 			`gorm:"primary_key;auto_increment"`
	FirstName string 		`gorm:"not null"`
	LastName  string 		`gorm:"not null"`
	Username  string 		`gorm:"not null;unique"`
	Email     *string 		`gorm:"not null;unique"`
	CreatedAt time.Time  	`gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time 	`gorm:"default:CURRENT_TIMESTAMP"`
}


type