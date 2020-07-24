package dao

import (
	"gold_hill/scaffold/model"
	"xorm.io/xorm"
)

type Attachments struct {
	Base
}

func NewAttachmentsDao() *Attachments {
	return &Attachments{}
}

func (d *Attachments) WithSession(s *xorm.Session) *Attachments {
	if s != nil {
		d.Write(s)
	} else {
		d.NewSession()
	}
	return d
}

func (d *Attachments) GetBy(sha1 string) (*model.Attachments, error) {
	var (
		m = &model.Attachments{}
	)
	_, err := d.session.Where("sha1=?", sha1).Get(m)
	if err != nil {
		return m, err
	}
	return m, nil
}
