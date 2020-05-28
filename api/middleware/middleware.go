package middleware

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/logging"
)

// Grant - CREATE A LITERAL OF THIS AT YOUR PERIL
type Grant struct {
	Claims auth.UserClaims
	// Convenience fields for checking auth
	IsAdmin    bool
	IsManager  bool
	IsDelegate bool
	// contains the sentry hub
	Logger logging.Logger
}

// Authenticate is used to verify and get access to middleware functions
func Authenticate(jwt string) (*Grant, error) {
	claims, err := auth.ValidateToken(jwt)
	if err != nil {
		glog.Infof("Authentication failed: %s", err.Error())
		return &Grant{}, &errors.ErrTokenInvalid
	}

	var (
		isAdmin    bool
		isManager  bool
		isDelegate bool
	)

	switch claims.Role {
	case auth.AdminRole:
		isAdmin = true
	case auth.ManagerRole:
		isManager = true
	case auth.DelegateRole:
		isDelegate = true
	}

	return &Grant{
		Claims:     claims,
		IsAdmin:    isAdmin,
		IsManager:  isManager,
		IsDelegate: isDelegate,
	}, nil
}

// MaxPageLimit is the maximum amount of returned datapoints
const MaxPageLimit = int32(100)

// getPage adds limit and offset to a query
func getPage(query *gorm.DB, page *gentypes.Page) (*gorm.DB, int32, int32) {
	var (
		limit  = MaxPageLimit
		offset int32
	)
	query.Limit(MaxPageLimit)
	if page != nil {
		if page.Offset != nil {
			offset = *page.Offset
			query = query.Offset(offset)
		}
		if page.Limit != nil && *page.Limit <= MaxPageLimit {
			limit = *page.Limit
			query = query.Limit(limit)
		}
	}
	return query, limit, offset
}

/* getOrdering adds orderBy to a query,

In no circumstances is "allowedFields" to be given by the user
*/
func getOrdering(query *gorm.DB, orderBy *gentypes.OrderBy, allowedFields []string, defaultOrdering string) (*gorm.DB, error) {
	if orderBy != nil {
		var allowed bool
		for _, field := range allowedFields {
			if orderBy.Field == field {
				allowed = true
				break
			}
		}

		if !allowed {
			glog.Infof("Ordering unauthorized: %s", orderBy.Field)
			return query, &errors.ErrOrderUnauthorized
		}

		ordering := "DESC"
		if orderBy.Ascending != nil && *orderBy.Ascending {
			ordering = "ASC"
		}
		// fmt.Sprintf is fine here as fields are checked against allowed ones.
		query = query.Order(fmt.Sprintf("%s %s", orderBy.Field, ordering))
	}

	// should order by the chosen field first, and then the default field
	query = query.Order(defaultOrdering)
	return query, nil
}

func filterUser(query *gorm.DB, filter *gentypes.UserFilter) *gorm.DB {
	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query = query.Where("first_name || ' ' || last_name ILIKE ?", "%%"+*filter.Name+"%%")
		}
		if filter.UUID != nil && *filter.UUID != "" {
			query = query.Where("uuid = ?", *filter.UUID)
		}
		if filter.JobTitle != nil && *filter.JobTitle != "" {
			query = query.Where("job_title ILIKE ?", "%%"+*filter.JobTitle+"%%")
		}
		if filter.Telephone != nil && *filter.Telephone != "" {
			query = query.Where("job_title ILIKE ?", "%%"+*filter.Telephone+"%%")
		}
	}

	return query
}
