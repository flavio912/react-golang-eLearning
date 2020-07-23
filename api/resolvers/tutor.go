package resolvers

import "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

type TutorResolver struct {
	Tutor gentypes.Tutor
}

func (t *TutorResolver) UUID() gentypes.UUID  { return t.Tutor.UUID }
func (t *TutorResolver) Name() string         { return t.Tutor.Name }
func (t *TutorResolver) CIN() string          { return t.Tutor.CIN }
func (t *TutorResolver) SignatureURL() string { return t.Tutor.SignatureURL }
