package repository

type HeadObject struct {
	FileName      string
	ContentLength int64
	ContentType   string
	LastModified  int64
}
type R2Repository interface {
	UploadFile(fileData *[]byte) (fileName string, err error)
	UploadBase64File(base64File *string) (fileName string, err error)
	HeadObject(fileName string) *HeadObject
	DeleteFile(fileName string) (err error)
}
