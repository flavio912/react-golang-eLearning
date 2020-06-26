package domain

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/users"
)

// This isn't really a DDD app, but this package name still made sense

type Application struct {
	CourseApp course.CourseApp
	UsersApp  users.UsersApp
	AdminApp  application.AdminApp
}
