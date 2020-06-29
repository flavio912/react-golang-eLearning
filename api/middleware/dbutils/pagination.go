package dbutils

import (
	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/middleware"
)

func (d *dbUtilsImpl) GetPageOf(
	model interface{},
	out interface{},
	page *gentypes.Page,
	orderBy *gentypes.OrderBy,
	allowedOrderBy []string,
	defaultOrderQuery string,
	filterFunc func(*gorm.DB) *gorm.DB,
) (gentypes.PageInfo, error) {
	// Count the total filtered dataset
	var count int32
	query := filterFunc(database.GormDB)
	countErr := query.Model(model).Limit(middleware.MaxPageLimit).Offset(0).Count(&count).Error
	if countErr != nil {
		d.Logger.Log(sentry.LevelError, countErr, "Unable to count")
		return gentypes.PageInfo{}, countErr
	}

	query, orderErr := middleware.GetOrdering(query, orderBy, allowedOrderBy, defaultOrderQuery)
	if orderErr != nil {
		d.Logger.Log(sentry.LevelError, countErr, "Unable to order")
		return gentypes.PageInfo{}, orderErr
	}

	query, limit, offset := middleware.GetPage(query, page)

	query = query.Find(out)
	if query.Error != nil {
		if query.RecordNotFound() {
			return gentypes.PageInfo{}, &errors.ErrNotFound
		}

		d.Logger.Log(sentry.LevelError, query.Error, "Unable to find")
		return gentypes.PageInfo{}, &errors.ErrWhileHandling
	}

	return gentypes.PageInfo{
		Total:  count,
		Offset: offset,
		Limit:  limit,
		Given:  0,
	}, nil
}
