package main

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID        uint
	FirstName string
	LastName  string
	Username  string
	Email     *string
	CreatedAt string
	UpdatedAt string
}
