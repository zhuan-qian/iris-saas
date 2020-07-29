package admins

import (
	"gold_hill/mine/app/controllers"
	"gold_hill/mine/service"
)

type Attachments struct {
	controllers.Base
}

func (c *Attachments) Post() {
	var (
		file, info, err = c.Ctx.FormFile("file")
		path            string
	)

	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	defer file.Close()

	path, err = service.NewAttachmentsService().FileToCloud(&file, info)

	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}
	c.SendSmile(path)
	return
}

func (c *Attachments) PostBy(way string) {
	var (
		file, info, err = c.Ctx.FormFile("file")
		path            string
	)

	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	defer file.Close()

	path, err = service.NewAttachmentsService().CourseToCloud(way, &file, info)

	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}
	c.SendSmile(path)
	return

}
