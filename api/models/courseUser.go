package models

type UserCourseInfo struct {
	User
	DelegateID *uint
	Delegate Delegate
	IndividualID *uint 
	Individual Individual
	// Progress, Assigned Courses etc
}