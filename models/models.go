package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	// ... tambahkan field lain yang diperlukan
}

type Campaign struct {
	gorm.Model
	Title       string
	Description string
	UserID      uint // Foreign Key ke User
	// ... tambahkan field lain yang diperlukan
}

type Donation struct {
	gorm.Model
	Amount     float64
	UserID     uint // Foreign Key ke User
	CampaignID uint // Foreign Key ke Campaign
	// ... tambahkan field lain yang diperlukan
}
