package middleware

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

// Grant - CREATE A LITERAL OF THIS AT YOUR PERIL
type Grant struct {
	Claims auth.UserClaims
	// Convenience functions for checking auth
	IsAdmin    bool
	IsManager  bool
	IsDelegate bool
}

// Authenticate is used to verify and get access to middleware functions
func Authenticate(jwt string) (*Grant, error) {
	claims, err := auth.ValidateToken(jwt)
	if err != nil {
		glog.Info(err.Error())
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
const MaxPageLimit = int32(400)

func getPage(query *gorm.DB, page *gentypes.Page) *gorm.DB {
	query.Limit(MaxPageLimit)
	if page != nil {
		if page.Limit != nil && *page.Limit <= MaxPageLimit {
			query = query.Limit(*page.Limit)
		}
		if page.Offset != nil {
			query = query.Offset(*page.Offset)
		}
	}
	return query
}
