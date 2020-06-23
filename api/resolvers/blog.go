package resolvers

import (
	"context"

	"github.com/golang/glog"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
)

type BlogResolver struct {
	Blog gentypes.Blog
}

func (b *BlogResolver) UUID() gentypes.UUID    { return b.Blog.UUID }
func (b *BlogResolver) Title() string          { return b.Blog.Title }
func (b *BlogResolver) Body() string           { return b.Blog.Body }
func (b *BlogResolver) HeaderImageURL() string { return b.Blog.HeaderImageURL }

// TODO: Use dataloaders
func (b *BlogResolver) Category() *CategoryResolver {
	return &CategoryResolver{
		Category: b.Blog.Category,
	}
}
func (b *BlogResolver) Author() *BlogAuthorResolver {
	return &BlogAuthorResolver{
		BlogAuthor: b.Blog.Author,
	}
}
func (b *BlogResolver) BlogBodyImages(ctx context.Context) (*[]*BlogImageResolver, error) {
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return nil, &errors.ErrUnauthorized
	}

	courseApp := course.NewCourseApp(grant)
	imgs, err := courseApp.GetBlogBodyImages(b.Blog.UUID)
	if err != nil {
		glog.Info("Unable to resolve blog body images")
		return nil, err
	}

	var res []*BlogImageResolver
	for _, img := range imgs {
		res = append(res, &BlogImageResolver{
			BlogImage: img,
		})
	}
	return &res, nil
}

type BlogAuthorResolver struct {
	BlogAuthor gentypes.BlogAuthor
}

func (ba *BlogAuthorResolver) FirstName() string      { return ba.BlogAuthor.FirstName }
func (ba *BlogAuthorResolver) LastName() string       { return ba.BlogAuthor.LastName }
func (ba *BlogAuthorResolver) ProfilePicURL() *string { return ba.BlogAuthor.ProfilePicURL }

type BlogImageResolver struct {
	BlogImage gentypes.BlogImage
}

func (bi *BlogImageResolver) JsonID() string { return bi.BlogImage.JsonID }
func (bi *BlogImageResolver) Url() string    { return bi.BlogImage.Url }
