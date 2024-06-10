package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:varchar(100);unique"`
	Password string `json:"password" gorm:"size:256"`
	Token    string `json:"token" gorm:"size:512"`
	Logged   bool   `json:"logged" gorm:"default:false"`
	Permessi string `json:"permission" gorm:"type:varchar(20)"`
}

func IsValidPermessi(permesso string) bool {
	validPermessi := []string{"Developers", "Admin", "Superuser", "Responsible", "User"}
	for _, v := range validPermessi {
		if permesso == v {
			return true
		}
	}
	return false
}
