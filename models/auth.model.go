package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	firstName   string
	lastName    string
	email       string
	password    string
	phoneNumber int8
	address     string
}
