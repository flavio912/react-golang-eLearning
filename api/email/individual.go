package email

func SendAccountCompleteEmail(fName string, email string) error {
	data := struct {
		FName string
	}{
		FName: fName,
	}

	return SendEmail("individualAccountCreated.html", "Welcome to TTC", email, data)
}
