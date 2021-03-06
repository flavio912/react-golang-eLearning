package resolvers

import (
	"context"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type CourseResolver struct {
	Course gentypes.Course
}

type NewCourseArgs struct {
	ID     *uint
	Course *gentypes.Course
}

func NewCourseResolver(ctx context.Context, args NewCourseArgs) (*CourseResolver, error) {
	if args.ID != nil {
		// TODO: Use loader to stop n+1 calls
		app := auth.AppFromContext(ctx)
		info, err := app.CourseApp.Course(*args.ID)
		if err != nil {
			return &CourseResolver{}, err
		}

		return &CourseResolver{
			Course: info,
		}, nil
	}
	if args.Course != nil {
		return &CourseResolver{Course: *args.Course}, nil
	}

	return &CourseResolver{}, &errors.ErrUnableToResolve
}

func (r *CourseResolver) Id() int32 {
	var item = int32(r.Course.ID)
	return item
}
func (r *CourseResolver) Type() gentypes.CourseType { return r.Course.CourseType }
func (r *CourseResolver) Name() string              { return r.Course.Name }
func (r *CourseResolver) BackgroundCheck() *bool {
	return helpers.BoolPointer(r.Course.BackgroundCheck)
}
func (r *CourseResolver) AccessType() gentypes.AccessType { return r.Course.AccessType }
func (r *CourseResolver) Price() float64                  { return r.Course.Price }
func (r *CourseResolver) Color() *string                  { return helpers.StringPointer(r.Course.Color) }
func (r *CourseResolver) Excerpt() *string                { return helpers.StringPointer(r.Course.Excerpt) }
func (r *CourseResolver) Introduction() *string {
	return helpers.StringPointer(r.Course.Introduction)
}
func (r *CourseResolver) HowToComplete() *string {
	return helpers.StringPointer(r.Course.HowToComplete)
}
func (r *CourseResolver) HoursToComplete() *float64 {
	return helpers.FloatPointer(r.Course.HoursToComplete)
}
func (r *CourseResolver) WhatYouLearn() *[]string {
	var learn = r.Course.WhatYouLearn
	return &learn
}
func (r *CourseResolver) Requirements() *[]string {
	var req = r.Course.Requirements
	return &req
}

func (r *CourseResolver) SpecificTerms() *string {
	return helpers.StringPointer(r.Course.SpecificTerms)
}
func (r *CourseResolver) Category(ctx context.Context) (*CategoryResolver, error) {
	if r.Course.CategoryUUID != nil {
		return NewCategoryResolver(ctx, NewCategoryResolverArgs{UUID: *r.Course.CategoryUUID})
	}
	return nil, nil
}
func (r *CourseResolver) AllowedToBuy() *bool {
	return helpers.BoolPointer(r.Course.AllowedToBuy)
}
func (r *CourseResolver) Syllabus(ctx context.Context) (*[]*SyllabusResolver, error) {
	return NewSyllabusResolvers(ctx, NewSyllabusArgs{CourseID: &r.Course.ID})
}
func (r *CourseResolver) BannerImageURL() *string {
	return r.Course.BannerImageURL
}
func (r *CourseResolver) ExpiresInMonths() int32 {
	return int32(r.Course.ExpiresInMonths)
}
func (r *CourseResolver) ExpirationToEndMonth() bool {
	return r.Course.ExpirationToEndMonth
}
func (r *CourseResolver) Published() bool {
	return r.Course.Published
}
func (r *CourseResolver) CertificateType(ctx context.Context) (*CertificateTypeResolver, error) {
	if r.Course.CertificateTypeUUID == nil {
		return nil, nil
	}

	return NewCertificateTypeResolver(ctx, NewCertificateTypeArgs{
		CertificateTypeUUID: r.Course.CertificateTypeUUID,
	})
}

type CoursePageResolver struct {
	edges    *[]*CourseResolver
	pageInfo *PageInfoResolver
}

func (r *CoursePageResolver) PageInfo() *PageInfoResolver { return r.pageInfo }
func (r *CoursePageResolver) Edges() *[]*CourseResolver   { return r.edges }
