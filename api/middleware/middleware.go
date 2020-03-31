package middleware

import (
	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/auth"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
)

// grant is purposely not exported so that the Authenticate method cannot be bypassed
type grant struct {
	Claims auth.UserClaims
	// Convenience functions for checking auth
	IsAdmin    bool
	IsManager  bool
	IsDelegate bool
}

// Authenticate is used to verify and get access to middleware functions
func Authenticate(jwt string) (*grant, error) {
	claims, err := auth.ValidateToken(jwt)
	if err != nil {
		glog.Info(err.Error())
		return &grant{}, &errors.ErrTokenInvalid
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

	return &grant{
		Claims:     claims,
		IsAdmin:    isAdmin,
		IsManager:  isManager,
		IsDelegate: isDelegate,
	}, nil
}
