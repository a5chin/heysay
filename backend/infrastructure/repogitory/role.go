package repogitory

import (
	"context"
	"errors"
	"heysay/entity"
	"heysay/infrastructure/driver"
	"heysay/infrastructure/repogitory/model"

	"net/http"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type RoleRepogitory struct {
	*driver.TokenDriver
}

func NewRoleRepogitory(tokenDriver *driver.TokenDriver) *RoleRepogitory {
	return &RoleRepogitory{tokenDriver}
}

func (r RoleRepogitory) CreateRole(ctx context.Context, roleName string) error {
	db, _ := ctx.Value(driver.TxKey).(*gorm.DB)
	if err := db.Create(
		&model.Role{
			ID:   model.GenerateID().String(),
			Name: roleName,
		},
	).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == driver.ErrDuplicateEntryNumber {
			return entity.WrapError(http.StatusConflict, err)
		}
		return err
	}
	return nil
}
