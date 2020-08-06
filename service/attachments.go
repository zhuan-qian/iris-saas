package service

import (
	"errors"
	"mime/multipart"
	"os"
	"zhuan-qian/go-saas/common"
	"zhuan-qian/go-saas/dao"
	"zhuan-qian/go-saas/model"
	"zhuan-qian/go-saas/service/file_cloud"
	"strings"
	"time"
)

type Attachments struct {
	d *dao.Attachments
}

func NewAttachmentsService() *Attachments {
	return &Attachments{d: dao.NewAttachmentsDao().WithSession(nil)}
}

func (s *Attachments) CourseToCloud(way string, file *multipart.File, info *multipart.FileHeader) (path string, err error) {
	var (
		o file_cloud.Oss
		f = &model.Attachments{}

		hashname          string
		filename          string
		mime              string
		mimeSplit         []string
		timeStr           string
		fileLimit              = ""
		mimeLimitPosition byte = 0
		putPrefix              = ""
		lookPrefix             = ""
		indexFile              = ""
	)

	//当前只用腾讯cos
	o = file_cloud.NewTencentOss()

	//sha1值计算
	hashname, err = common.GenerateSha1ByFile(file)
	if err != nil {
		return "", err
	}

	f, err = s.d.GetBy(hashname)
	if err != nil {
		return "", err
	}
	if f.Id != 0 {
		return f.Path, err
	}

	//格式校验
	mime = info.Header.Get("Content-Type")
	mimeSplit = strings.Split(mime, "/")
	if len(mimeSplit) != 2 {
		err = errors.New("文件类型未知")
		return "", err
	}

	//处理方式
	switch way {
	case "h5pack":
		fileLimit = "zip"
		putPrefix = "pack_courses/"
		lookPrefix = "courses/h5"
		mimeLimitPosition = 1
		indexFile = "story_html5.html"

	case "video":
		fileLimit = "video"

	default:
		return "", errors.New("invalid way")
	}

	//文件类型校验
	if !strings.Contains(fileLimit, mimeSplit[mimeLimitPosition]) {
		err = errors.New("不支持的文件类型")
		return "", err
	}

	//配置存储路径
	timeStr = time.Now().Format("20060102")
	filename = hashname + "." + s.getFileNameBySubMime(mimeSplit[1])
	f.Path = info.Header.Get("Content-Type") + "/" + putPrefix + timeStr + "/" + filename

	err = o.UploadFile(f.Path, file, info)
	if err != nil {
		return "", err
	}

	if lookPrefix != "" {
		f.Path = file_cloud.GetCosByBucket(os.Getenv("COS_UPLOAD_BUCKET")) + "/" + lookPrefix + "/" +
			hashname + "/" + indexFile
	} else {
		f.Path = file_cloud.GetCosByBucket(os.Getenv("COS_UPLOAD_BUCKET")) + "/" + f.Path
	}

	f.Size = info.Size
	f.Sha1 = hashname
	f.Mime = info.Header.Get("Content-Type")
	f.Status = 1

	_, err = s.d.InsertOne(f)
	if err != nil {
		return "", err
	}

	return f.Path, nil
}

func (s *Attachments) FileToCloud(file *multipart.File, info *multipart.FileHeader) (string, error) {
	var (
		o file_cloud.Oss
		f = &model.Attachments{}

		hashname  string
		filename  string
		mime      string
		mimeSplit []string
		fileLimit = "image,video,audio"
		timeStr   string
		err       error
	)

	//当前只用腾讯cos
	o = file_cloud.NewTencentOss()

	//sha1值计算
	hashname, err = common.GenerateSha1ByFile(file)
	if err != nil {
		return "", err
	}

	f, err = s.d.GetBy(hashname)
	if err != nil {
		return "", err
	}
	if f.Id != 0 {
		return f.Path, err
	}

	//格式校验
	mime = info.Header.Get("Content-Type")
	mimeSplit = strings.Split(mime, "/")
	if len(mimeSplit) != 2 {
		err = errors.New("文件类型未知")
		return "", err
	}

	//文件类型校验
	if !strings.Contains(fileLimit, mimeSplit[0]) {
		err = errors.New("不支持的文件类型")
		return "", err
	}

	timeStr = time.Now().Format("20060102")
	filename = hashname + "." + s.getFileNameBySubMime(mimeSplit[1])

	f.Path = info.Header.Get("Content-Type") + "/" + timeStr + "/" + filename

	err = o.UploadFile(f.Path, file, info)
	if err != nil {
		return "", err
	}
	f.Path = file_cloud.GetCosByBucket(os.Getenv("COS_UPLOAD_BUCKET")) + "/" + f.Path
	f.Size = info.Size
	f.Sha1 = hashname
	f.Mime = info.Header.Get("Content-Type")
	f.Status = 1

	_, err = s.d.InsertOne(f)
	if err != nil {
		return "", err
	}

	return f.Path, nil
}

func (s *Attachments) getFileNameBySubMime(m string) string {
	switch m {
	case "jpeg":
		fallthrough
	case "jpg":
		fallthrough
	case "png":
		return "jpg"
	default:
		return m
	}
}
