package models

type User struct {
	Basc
	Name        string `json:"name"  gorm:"type:varchar(100)"`
	Email       string `json:"email" gorm:"uniqueIndex;not null"`
	PhoneNumber string `json:"phoneNumber"  gorm:"type:varchar(15)"`
	Password    string `json:"-" gorm:"not null"`
	Status      bool   `json:"status" gorm:"default:true"`
}
