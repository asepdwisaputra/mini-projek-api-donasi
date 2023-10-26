package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `gorm:"primary_key" json:"id"`
	Name      string         `json:"name" form:"name"`
	Telephone string         `json:"telephone" form:"telephone"`
	Email     string         `json:"email" form:"email"`
	Password  string         `json:"password" form:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at"`

	// Hubungan User dengan Campaign (One-to-Many)
	//Campaigns []Campaign `gorm:"foreignkey:UserID"`

	// Hubungan User dengan Donation (One-to-Many)
	//Donations []Donation `gorm:"foreignkey:UserID"`
}

type UserResponseJWT struct {
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}

// type UserResponse struct {
// 	ID        int    `json:"id" form:"id"`
// 	Name      string `json:"name" form:"name"`
// 	Email     string `json:"email" form:"email"`
// 	Telephone string    `json:"telephone" form:"telephone"`
// }

type Campaign struct {
	ID             int       `json:"id" form:"id"`
	Title          string    `json:"title" form:"title"`
	Description    string    `json:"description" form:"description"`
	Start          time.Time `json:"start" form:"start"`
	End            time.Time `json:"end" form:"end"`
	CPocket        string    `json:"cpocket" form:"cpocket"`
	Status         string    `json:"status" form:"status"`
	Photo          string    `json:"photo" form:"photo"`
	TotalCollected float64   `json:"total_collected" form:"total_collected"`

	// Foreign Key ke User
	UserID int `json:"user_id"`
	//ABAIKAN INI--gorm:"foreignkey:ID"

	// Hubungan Campaign dengan User
	User User `gorm:"foreignkey:UserID" json:"user"` // Juga mendefinisikan hubungan Campaign dengan User

	// // Hubungan Campaign dengan Donation (One-to-Many)
	// Donations []Donation `gorm:"foreignkey:CampaignID"`
}

type Donation struct {
	ID     int       `json:"id" form:"id"`
	Amount float64   `json:"amount" form:"amount"`
	Date   time.Time `json:"date" form:"date"`
	Status string    `json:"status" form:"status"`

	// Foreign Key ke User
	UserID int `json:"user_id" gorm:"foreignkey:ID"`
	// Foreign Key ke Campaign
	CampaignID int `json:"campaign_id" gorm:"foreignkey:ID"`

	// // Sebuah hook Gorm sebelum membuat donasi
	// BeforeCreate(scope *gorm.Scope) error {
	//     // Mengambil nilai Amount dari donasi
	//     amount, _ := scope.FieldByName("Amount")

	//     // Mengambil ID kampanye dari donasi
	//     campaignID, _ := scope.FieldByName("CampaignID")

	//     // Menghitung total terkumpul sebelum donasi
	//     var totalCollected float64
	//     scope.DB().Model(&Campaign{}).Where("id = ?", campaignID.Field.Interface()).Pluck("total_collected", &totalCollected)

	//     // Menambah nilai Amount ke total terkumpul
	//     totalCollected += amount.Field.Interface().(float64)

	//     // Mengupdate total terkumpul dalam kampanye
	//     scope.DB().Model(&Campaign{}).Where("id = ?", campaignID.Field.Interface()).Update("total_collected", totalCollected)

	//     return nil
	// }
}
