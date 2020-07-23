package file_cloud

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/denverdino/aliyungo/sts"
	"mime/multipart"
	"os"
)

type aliOss struct {
	c          *oss.Client
	bucketName string
}

func NewAliOss() (*aliOss, error) {
	// 创建OSSClient实例。
	client, err := oss.New(os.Getenv("OSS_ENDPOINT"), os.Getenv("ALI_KEY_ID"), os.Getenv("ALI_KEY_SECRET"))
	return &aliOss{c: client, bucketName: os.Getenv("OSS_BUCKET")}, err
}

func (s *aliOss) newBucket() (*oss.Bucket, error) {
	return s.c.Bucket(s.bucketName)
}

func (s *aliOss) UploadFile(filePath string, file *multipart.File, info *multipart.FileHeader) error {
	b, err := s.newBucket()
	if err != nil {
		return err
	}

	//上传至oss
	err = b.PutObject(filePath, *file, oss.ContentType(info.Header.Get("Content-Type")+";"))
	return err
}

//func (s *aliOss) SetAccess(filePath string, access string) bool {
//	b, err := s.newBucket()
//	if err != nil {
//		return false
//	}
//	if access == model.FILE_ACCESS_IS_PUBLIC {
//		err = b.SetObjectACL(filePath, oss.ACLPublicRead)
//	} else {
//		err = b.SetObjectACL(filePath, oss.ACLPrivate)
//	}
//	if err != nil {
//		return false
//	}
//	return true
//}

func GetSts() (*sts.AssumeRoleResponse, error) {
	client := sts.NewClientWithEndpoint(os.Getenv("STS_ENDPOINT"), os.Getenv("ALI_KEY_ID"), os.Getenv("ALI_KEY_SECRET"))
	request := sts.AssumeRoleRequest{RoleArn: os.Getenv("STS_RAM"), RoleSessionName: "alice", DurationSeconds: 3600, Policy: ""}
	a, err := client.AssumeRole(request)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
