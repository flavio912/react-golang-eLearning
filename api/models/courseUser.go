package models

type UserCourseInfo struct {
	DelegateID   *uint
	Delegate     Delegate
	IndividualID *uint
	Individual   Individual
	// Progress, Assigned Courses etc
}
