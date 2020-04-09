package gentypes

type UploadFileMeta struct {
	FileType      string
	ContentLength int32
}

type UploadFileResp struct {
	URL          string
	SuccessToken string
}

type UploadFileSuccess struct {
	SuccessToken string
}

type Address struct {
	ID           uint
	AddressLine1 string
	AddressLine2 string
	County       string
	PostCode     string
	Country      string
}
