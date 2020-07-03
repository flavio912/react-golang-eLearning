package dbutils

import (
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
)

type DbUtils interface {
	GetPageOf(
		model interface{},
		out interface{},
		page *gentypes.Page,
		orderBy *gentypes.OrderBy,
		allowedOrderBy []string,
		defaultOrderQuery string,
		filterFunc func(*gorm.DB) *gorm.DB,
	) (gentypes.PageInfo, error)
}

type dbUtilsImpl struct {
	Logger *logging.Logger
}

func NewDBUtils(logger *logging.Logger) DbUtils {
	return &dbUtilsImpl{
		Logger: logger,
	}
}
