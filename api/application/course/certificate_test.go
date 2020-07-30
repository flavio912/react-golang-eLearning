package course_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/application/course"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

func TestGeneratePdfFromUrl(t *testing.T) {
	assert.Nil(t, helpers.LoadConfig())

	_, err := course.GeneratePdfFromURL("https://google.com")
	assert.Nil(t, err)
}
