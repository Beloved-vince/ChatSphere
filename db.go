package main

import "github.com/jinzhu/gorm"

var db *gorm.DB

func ConnectDatabase() {
	connectionString := "host=localhost port=5432 user=postgres dnname=postgres password=Oludare2001 sslmode=disable"
	var err error

	db, err = gorm.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}, &Room{}, Message{})
}
