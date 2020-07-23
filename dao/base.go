package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"scaffold/common"
	"xorm.io/xorm"
)

type BaseDao interface {
	NewSession()
	Write(session *xorm.Session)
	Read() *xorm.Session
	Insert(bean interface{}) (int64, error)
	InsertMulti(rows interface{}) (int64, error)
}

type Base struct {
	session *xorm.Session
}

func (b *Base) NewSession() {
	b.session = common.GetDB().NewSession()
}

func (b *Base) Write(session *xorm.Session) {
	b.session = session
}

func (b *Base) Read() *xorm.Session {
	return b.session
}

func (b *Base) InsertOne(bean interface{}) (int64, error) {
	return b.session.InsertOne(bean)
}

func (b *Base) Insert(bean interface{}) (int64, error) {
	return b.session.Insert(bean)
}
func (b *Base) InsertMulti(rows interface{}) (int64, error) {
	return b.session.InsertMulti(rows)
}

func (b *Base) Find(rowsSlicePtr interface{}, condiBean ...interface{}) error {
	return b.session.Find(rowsSlicePtr, condiBean...)
}
func (b *Base) FindAndCount(rowsSlicePtr interface{}, condiBean ...interface{}) (int64, error) {
	return b.session.FindAndCount(rowsSlicePtr, condiBean...)
}

func (b *Base) InfoBy(cols []string, id int64, bean interface{}) (bool, error) {
	if cols != nil {
		b.session = b.session.Cols(cols...)
	}
	return b.session.ID(id).Get(bean)
}

func (b *Base) ModifyBy(cols []string, id int64, bean interface{}) (int64, error) {
	return b.session.Cols(cols...).ID(id).Update(bean)
}

func (b *Base) ModifyByField(cols []string, field string, value interface{}, bean interface{}) (int64, error) {
	return b.session.Cols(cols...).Where(field+"=?", value).Update(bean)
}

func (b *Base) DeleteBy(table string, id int64) error {
	sql := fmt.Sprintf("delete from `%s` where id=%d", table, id)
	_, err := b.session.Exec(sql)
	return err
}

func (b *Base) DeleteByField(table string, field string, value string) error {
	sql := fmt.Sprintf("delete from `%s` where %s='%s'", table, field, value)
	_, err := b.session.Exec(sql)
	return err
}

func (b *Base) ExistByField(field string, value interface{}, bean interface{}) (bool, error) {
	return b.session.Where(field+"=?", value).Exist(bean)
}
