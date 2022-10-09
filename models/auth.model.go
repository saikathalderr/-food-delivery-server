package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint           `gorm:"primaryKey"`
	FirstName   string         `json:"firstName"`
	LastName    string         `json:"lastName"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	PhoneNumber int            `json:"phoneNumber"`
	Address     string         `json:"address"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type MenuItem struct {
	gorm.Model
	ID           uint           `gorm:"primaryKey"`
	Name         string         `json:"name"`
	Price        string         `json:"price"`
	SpecialPrice string         `json:"specialPrice"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
type LocalBusiness struct {
	gorm.Model
	ID               uint           `gorm:"primaryKey"`
	OrganizationName string         `json:"organizationName"`
	Email            string         `json:"email"`
	Password         string         `json:"password"`
	Address          string         `json:"address"`
	PhoneNumber      int            `json:"phoneNumber"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Menu             []MenuItem     `json:"menu,omitempty" gorm:"foreignKey:ID;references:ID"`
}
