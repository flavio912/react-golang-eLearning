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
