package loader

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/errors"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler/auth"
)

type blogLoader struct {
}

func sortBlogs(blogs []gentypes.Blog, keys dataloader.Keys) []gentypes.Blog {
	var (
		k       = keys.Keys()
		blogMap = map[string]gentypes.Blog{}
		sorted  = make([]gentypes.Blog, len(k))
	)

	for _, lesson := range blogs {
		blogMap[lesson.UUID.String()] = lesson
	}

	for i, key := range keys {
		sorted[i] = blogMap[key.String()]
	}

	return sorted
}

func (b *blogLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	n := len(keys)

	// app := auth.AppFromContext(ctx)
	grant := auth.GrantFromContext(ctx)
	if grant == nil {
		return loadBatchError(&errors.ErrUnauthorized, n)
	}

	blogApp := application.NewBlogApp(grant)
	blogs, err := blogApp.GetBlogsByUUID(keys.Keys())
	if err != nil {
		return loadBatchError(err, n)
	}

	blogs = sortBlogs(blogs, keys)
	res := make([]*dataloader.Result, n)
	for i, blog := range blogs {
		res[i] = &dataloader.Result{Data: blog}
	}
	return res
}

func LoadBlog(ctx context.Context, uuid string) (gentypes.Blog, error) {
	var blog gentypes.Blog
	data, err := extractAndLoad(ctx, blogLoaderKey, uuid)
	if err != nil {
		return blog, err
	}

	blog, ok := data.(gentypes.Blog)
	if !ok {
		return blog, fmt.Errorf("Wrong type: %T", data)
	}

	return blog, nil
}

type BlogResult struct {
	Blog  gentypes.Blog
	Error error
}

func LoadBlogs(ctx context.Context, uuids []string) ([]BlogResult, error) {
	ldr, err := extract(ctx, blogLoaderKey)
	if err != nil {
		return nil, err
	}

	data, errs := ldr.LoadMany(ctx, dataloader.NewKeysFromStrings(uuids))()

	results := make([]BlogResult, 0, len(uuids))

	for i, d := range data {
		var e error
		if errs != nil {
			e = errs[i]
		}

		blog, ok := d.(gentypes.Blog)
		if !ok && e == nil {
			e = fmt.Errorf("Wrong type: %T", blog)
		}

		results = append(results, BlogResult{Blog: blog, Error: e})
	}

	return results, nil
}
