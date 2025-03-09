package repository

import (
	"github.com/Nurt0re/chatik"
	"gorm.io/gorm"
)

type UpdPostgres struct {
	db *gorm.DB
}

func NewUpdPostgres(db *gorm.DB) *UpdPostgres {
	return &UpdPostgres{
		db: db,
	}
}

func (r *UpdPostgres) UpdateUser(id int, input chatik.User) error {

	if err := r.db.Model(&chatik.User{}).Where("id = ?", id).Updates(input).Error; err != nil {
		return err
	}
	return nil
}

func (r *UpdPostgres) DeleteUser(id int) error {
	if err := r.db.Delete(&chatik.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
func (r *UpdPostgres) GetUser(id int) (chatik.User, error) {
	var user chatik.User
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}
func (r *UpdPostgres) GetAllUsers() ([]chatik.User, error) {
	var users []chatik.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}