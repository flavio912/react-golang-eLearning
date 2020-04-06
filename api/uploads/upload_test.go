package uploads

import (
	"fmt"
	"testing"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
)

func TestProfileUpload(t *testing.T) {
	helpers.LoadConfig("../config.yml")
	url, _ := generateProfileURL("jpg")
	fmt.Printf("URL: %s \n", url)
}
