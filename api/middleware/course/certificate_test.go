package course_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

func TestCreateCertificateType(t *testing.T) {
	t.Run("Creates a certificate type", func(t *testing.T) {
		inp := gentypes.CreateCertificateTypeInput{
			Name:           "Rap God",
			RegulationText: "Good at rapping",
		}

		certType, err := courseRepo.CreateCertificateType(inp)

		assert.Nil(t, err)

		assert.Equal(t, inp.Name, certType.Name)
		assert.Equal(t, inp.RegulationText, certType.RegulationText)
	})
}

func TestCertificateTypes(t *testing.T) {
	prepareTestDatabase()

	t.Run("Should return all certificate types", func(t *testing.T) {
		certTypes, _, err := courseRepo.CertificateTypes(nil, nil)

		assert.Nil(t, err)
		assert.Len(t, certTypes, 3)
	})

	t.Run("Should page", func(t *testing.T) {
		limit := int32(2)
		page := gentypes.Page{Limit: &limit, Offset: nil}
		certTypes, pageInfo, err := courseRepo.CertificateTypes(&page, nil)

		assert.Nil(t, err)
		assert.Equal(t, gentypes.PageInfo{Total: 3, Given: 2, Limit: limit}, pageInfo)
		assert.Len(t, certTypes, 2)
	})

	tests := []struct {
		name    string
		filter  gentypes.CertificateTypeFilter
		wantLen int
	}{
		{
			name:    "name",
			filter:  gentypes.CertificateTypeFilter{Name: helpers.StringPointer("ava")},
			wantLen: 1,
		},
		{
			name:    "regulation text",
			filter:  gentypes.CertificateTypeFilter{RegulationText: helpers.StringPointer("tom")},
			wantLen: 1,
		},
		{
			name:    "CAANo requirement",
			filter:  gentypes.CertificateTypeFilter{RequiresCAANo: helpers.BoolPointer(true)},
			wantLen: 2,
		},
		{
			name:    "Showing Training Section",
			filter:  gentypes.CertificateTypeFilter{ShowTrainingSection: helpers.BoolPointer(true)},
			wantLen: 2,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Should filter by %s", test.name), func(t *testing.T) {
			certTypes, _, err := courseRepo.CertificateTypes(nil, &test.filter)

			assert.Nil(t, err)
			assert.Len(t, certTypes, test.wantLen)
		})
	}
}
