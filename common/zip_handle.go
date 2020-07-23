package common

import (
	"archive/zip"
	"errors"
	"io"
	"mime/multipart"
	"os"
)

type zipDecompressor struct {
	TargetDirector string
	File           *multipart.File
	FileHeader     *multipart.FileHeader
}

func NewZipDecompressor(zipFile *multipart.File, fileHeader *multipart.FileHeader, targetDirector string) *zipDecompressor {
	return &zipDecompressor{
		TargetDirector: targetDirector,
		File:           zipFile,
		FileHeader:     fileHeader,
	}
}

func (z *zipDecompressor) Do() (destDirectory string, err error) {
	var (
		closer   *zip.Reader
		sha1Name string
		exist    bool
	)

	if z.FileHeader == nil {
		return destDirectory, errors.New("未知文件类型,请确认文件是否正确")
	}

	if z.FileHeader.Header.Get("Content-Type") != "application/zip" {
		return destDirectory, errors.New("文件类型必须为zip压缩包")
	}

	closer, err = zip.NewReader(*z.File, z.FileHeader.Size)
	if err != nil {
		return destDirectory, errors.New("加载文件出错")
	}

	sha1Name, err = GenerateSha1ByFile(z.File)
	destDirectory = z.TargetDirector + "/" + sha1Name
	exist, err = PathExists(destDirectory)
	if err != nil {
		return
	}
	if !exist {
		if err = os.MkdirAll(destDirectory, 0755); err != nil {
			return
		}
	}

	for _, f := range closer.File {
		index := f.Name[:len(f.Name)-len(f.FileInfo().Name())]
		if index != "" {
			if err = os.MkdirAll(destDirectory+"/"+index, 0755); err != nil {
				return
			}
		}
		inFile, err := f.Open()
		if err != nil {
			return destDirectory, err
		}
		defer inFile.Close()

		outFile, err := os.OpenFile(destDirectory+"/"+f.Name, os.O_WRONLY|os.O_CREATE, 0754)
		if err != nil {
			return destDirectory, err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, inFile)
		if err != nil {
			return destDirectory, err
		}
	}
	return
}
