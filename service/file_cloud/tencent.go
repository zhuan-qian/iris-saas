package file_cloud

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

type tencentOss struct {
	c *cos.Client
}

func NewTencentOss() *tencentOss {
	o := &tencentOss{c: cos.NewClient(&cos.BaseURL{}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("TX_KEY_ID"),
			SecretKey: os.Getenv("TX_KEY_SECRET"),
		},
	})}
	o.SetBucket(GetCosByBucket(os.Getenv("COS_UPLOAD_BUCKET")))
	return o
}

func GetCosByBucket(bucket string) string {
	return fmt.Sprintf(os.Getenv("COS_ENDPOINT"), bucket)
}

func (s *tencentOss) SetBucket(bucketUrl string) *tencentOss {
	s.c.BaseURL.BucketURL, _ = url.Parse(bucketUrl)
	return s
}

//直接上传文件
func (s *tencentOss) UploadFile(targetPath string, file *multipart.File, info *multipart.FileHeader) error {
	_, err := s.c.Object.Put(context.Background(), targetPath, *file, &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{ContentType: info.Header.Get("Content-Type") + ";"},
	})
	return err
}

//通过本地磁盘上传对象
func (s *tencentOss) UploadFileFromPath(targetPath string, filePath string) error {
	_, err := s.c.Object.PutFromFile(context.Background(), targetPath, filePath, nil)
	return err
}
