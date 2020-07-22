package course_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"
)

func TestCreateCertificateType(t *testing.T){
	t.Run("Creates a certificate type", func (t *testing.T){
		inp := gentypes.CreateCertificateTypeInput{
			Name: "Rap God",
			RegulationText: "Good at rapping",
		}

		certType, err := courseRepo.CreateCertificateType(inp)

		assert.Nil(t, err)

		assert.Equal(t, inp.Name, certType.Name)
		assert.Equal(t, inp.RegulationText, certType.RegulationText)
	})
}