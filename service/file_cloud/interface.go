package file_cloud

import "mime/multipart"

type Oss interface {
	UploadFile(targetPath string, file *multipart.File, info *multipart.FileHeader) error
}
