package chatik

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Username string `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	Password string `gorm:"type:text;not null" json:"password"`
}
