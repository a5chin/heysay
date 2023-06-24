package model

import (
	"heysay/entity"

	"gorm.io/gorm"
)

type Role struct {
	ID   string
	Name string
	gorm.Model
}

func (m Role) ToEntity() *entity.Role {
	return &entity.Role{
		ID:   m.ID,
		Name: m.Name,
	}
}
