package repogitory

import (
	"context"
	"heysay/entity"
	"heysay/infrastructure/driver"
	"heysay/infrastructure/repogitory/model"

	"gorm.io/gorm"
)

type UserRepogitory struct {
	*driver.TokenDriver
}

func NewUserRepogitory(tokenDriver *driver.TokenDriver) *UserRepogitory {
	return &UserRepogitory{tokenDriver}
}

func (r UserRepogitory) GetUsers(ctx context.Context) ([]*entity.User, error) {
	records := []*model.User{}
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if err := db.Preload("Roles").Find(&records).Error; err != nil {
		return nil, err
	}
	var users []*entity.User
	for _, record := range records {
		users = append(users, record.ToEntity())
	}
	return users, nil
}

func (r UserRepogitory) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	var user *model.User
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if err := db.Preload("Roles").First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	return user.ToEntity(), nil
}
