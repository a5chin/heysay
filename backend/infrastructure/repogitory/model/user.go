package model

import (
	"heysay/entity"

	"gorm.io/gorm"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	gorm.Model
}

func (m User) ToEntity() *entity.User {
	return &entity.User{
		ID:    m.ID,
		Name:  m.Name,
		Email: m.Email,
	}
}
