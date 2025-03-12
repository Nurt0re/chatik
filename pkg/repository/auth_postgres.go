package repository

import (


	"github.com/Nurt0re/chatik"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(user chatik.User) (int, error) {
	var id int
	result:= r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	id = int(user.ID)
	return id, nil
}



func (r *AuthPostgres) GetUser(email, password string) (chatik.User, error) {
	var user chatik.User

	r.db.First(&user, "email = ? AND password = ?", email, password)
	return user, nil
}