package model

import (
	"heysay/entity"

	"gorm.io/gorm"
)

type User struct {
	ID       string
	Name     string
	Roles    []*Role `gorm:"many2many:user_roles"`
	Email    string
	Password string
	gorm.Model
}

func (m User) ToEntity() *entity.User {
	roles := []*entity.Role{}
	for _, r := range m.Roles {
		roles = append(roles, r.ToEntity())
	}
	return &entity.User{
		ID:    m.ID,
		Name:  m.Name,
		Roles: roles,
		Email: m.Email,
	}
}
