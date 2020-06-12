package models

type PendingOrder struct {
	Base
	StripeClientSecret string        `gorm:"unique"`
	Courses            []Course      `gorm:"many2many:pending_order_course_link;"`
	CourseTakers       []CourseTaker `gorm:"many2many:pending_order_course_takers;"`
	ExtraInvoiceEmail  string
}
