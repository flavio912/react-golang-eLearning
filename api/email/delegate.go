package email

func SendFinaliseAccountEmail(token string, fName string, email string) error {
	data := struct {
		Token string
		FName string
	}{
		Token: token,
		FName: fName,
	}

	return SendEmail("finaliseAccount.html", "TTC - Finalise Account", email, data)
}
