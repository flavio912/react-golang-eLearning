package resolvers

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
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

type BlogAuthorResolver struct {
	BlogAuthor gentypes.BlogAuthor
}

func (ba *BlogAuthorResolver) FirstName() string      { return ba.BlogAuthor.FirstName }
func (ba *BlogAuthorResolver) LastName() string       { return ba.BlogAuthor.LastName }
func (ba *BlogAuthorResolver) ProfilePicURL() *string { return ba.BlogAuthor.ProfilePicURL }
