package directory

import (
	"os"
	"gold_hill/mine/common"
)

func InitDirectory ()error{
	var(
		err error
	)

	err=publicResourceInit()
	if err!=nil{
		return err
	}
	return nil
}


//公共资源目录创建
func publicResourceInit() error{
	var(
		exist bool
		err error
	)

	exist,err=common.PathExists("public/resource")
	if err!=nil{
		return err
	}

	if !exist{
		err=os.Mkdir("public/resource",0755)
		return err
	}
	return nil
}
