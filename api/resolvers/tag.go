package resolvers

import (
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

type TagResolver struct {
	Tag gentypes.Tag
}

func (t *TagResolver) UUID() *gentypes.UUID { return helpers.UUIDPointer(t.Tag.UUID) }
func (t *TagResolver) Name() string         { return t.Tag.Name }
func (t *TagResolver) Color() string        { return t.Tag.Color }
