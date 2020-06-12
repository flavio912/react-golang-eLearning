package models

type PendingOrder struct {
	Base
	StripeClientSecret string        `gorm:"unique"`
	Courses            []Course      `gorm:"many2many:pending_order_course_link;association_autoupdate:false"`
	CourseTakers       []CourseTaker `gorm:"many2many:pending_order_course_takers;association_autoupdate:false"`
	ExtraInvoiceEmail  *string
}
